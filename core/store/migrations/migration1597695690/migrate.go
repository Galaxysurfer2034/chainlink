package migration1597695690

import (
	"github.com/jinzhu/gorm"
)

// Migrate creates the offchain_reporting_job_specs table
func Migrate(tx *gorm.DB) error {
	return tx.Exec(`
        CREATE TABLE pipeline_specs (
            id SERIAL PRIMARY KEY,
            dot_dag_source TEXT NOT NULL,
            created_at timestamptz NOT NULL
        );

        CREATE INDEX idx_pipeline_specs_created_at ON pipeline_specs USING BRIN (created_at);

        CREATE TABLE pipeline_task_specs (
            id SERIAL PRIMARY KEY,
            dot_id TEXT,
            pipeline_spec_id INT NOT NULL REFERENCES pipeline_specs (id) ON DELETE CASCADE,
            type TEXT NOT NULL,
            json jsonb NOT NULL,
            index INT NOT NULL DEFAULT 0,
            successor_id INT REFERENCES pipeline_task_specs (id),
            created_at timestamptz NOT NULL
        );

        CREATE INDEX idx_pipeline_task_specs_pipeline_spec_id ON pipeline_task_specs (pipeline_spec_id);
        CREATE INDEX idx_pipeline_task_specs_successor_id ON pipeline_task_specs (successor_id);
        CREATE INDEX idx_pipeline_task_specs_created_at ON pipeline_task_specs USING BRIN (created_at);

        CREATE TABLE pipeline_runs (
            id BIGSERIAL PRIMARY KEY,
            pipeline_spec_id BIGINT NOT NULL REFERENCES pipeline_specs (id) ON DELETE CASCADE,
            meta jsonb,
            created_at timestamptz NOT NULL
            -- NOTE: Could denormalize here with finished_at/output/error of last task_run if that proves necessary for performance
        );

        CREATE INDEX idx_pipeline_runs_pipeline_spec_id ON pipeline_runs (pipeline_spec_id);
        CREATE INDEX idx_pipeline_runs_created_at ON pipeline_runs USING BRIN (created_at);

        CREATE TABLE pipeline_task_runs (
            id BIGSERIAL PRIMARY KEY,
            dot_id TEXT,
            pipeline_run_id BIGINT NOT NULL REFERENCES pipeline_runs (id) ON DELETE CASCADE,
            output JSONB,
            error TEXT,
            pipeline_task_spec_id INT NOT NULL REFERENCES pipeline_task_specs (id) ON DELETE CASCADE,
            index INT NOT NULL DEFAULT 0,
            created_at timestamptz NOT NULL,
            finished_at timestamptz,
            CONSTRAINT chk_pipeline_task_run_fsm CHECK (
                error IS NULL AND output IS NULL AND finished_at IS NULL
                OR
                error IS NULL AND output IS NOT NULL AND finished_at IS NOT NULL
                OR
                output IS NULL AND error IS NOT NULL AND finished_at IS NOT NULL
            )
        );


        ---
        --- Notify the Chainlink node when a new pipeline run has started
        ---

        CREATE OR REPLACE FUNCTION notifyPipelineRunStarted() RETURNS TRIGGER AS $_$
        BEGIN
            PERFORM pg_notify('pipeline_run_started', NEW.id::text);
            RETURN NEW;
        END
        $_$ LANGUAGE 'plpgsql';

        CREATE TRIGGER notify_pipeline_run_started
        AFTER INSERT ON pipeline_runs
        FOR EACH ROW EXECUTE PROCEDURE notifyPipelineRunStarted();


        ---
        --- Notify the Chainlink node when a pipeline run has completed
        ---

        CREATE OR REPLACE FUNCTION notifyPipelineRunCompleted() RETURNS TRIGGER AS $_$
        DECLARE done BOOLEAN;
        BEGIN
            SELECT bool_and(pipeline_task_runs.error IS NOT NULL OR pipeline_task_runs.output IS NOT NULL)
                INTO done
                FROM pipeline_runs
                JOIN pipeline_task_runs ON pipeline_task_runs.pipeline_run_id = pipeline_runs.id
                WHERE pipeline_runs.id = NEW.pipeline_run_id;

            IF done = TRUE THEN
                PERFORM pg_notify('pipeline_run_completed', NEW.pipeline_run_id::text);
            END IF;
            RETURN NEW;
        END
        $_$ LANGUAGE 'plpgsql';

        CREATE TRIGGER notify_pipeline_run_completed
        AFTER UPDATE ON pipeline_task_runs
        FOR EACH ROW EXECUTE PROCEDURE notifyPipelineRunCompleted();



        -- NOTE: This table is large and insert/update heavy so we must be efficient with indexes

        CREATE INDEX idx_pipeline_task_runs ON pipeline_task_runs USING BRIN (created_at);
        -- CREATE INDEX idx_pipeline_task_runs_finished_at ON pipeline_task_runs WHERE finished_at IS NOT NULL USING BRIN (finished_at);

        -- This query is used in the runner to find unstarted task runs
        -- CREATE INDEX idx_pipeline_task_runs_unfinished ON pipeline_task_runs (finished_at) WHERE finished_at IS NULL;

        ALTER TABLE offchainreporting_oracle_specs ADD COLUMN contract_address bytea NOT NULL;
        ALTER TABLE offchainreporting_oracle_specs ADD COLUMN p2p_peer_id text NOT NULL REFERENCES encrypted_p2p_keys (peer_id);
        ALTER TABLE offchainreporting_oracle_specs ADD COLUMN p2p_bootstrap_peers text[] NOT NULL;
        ALTER TABLE offchainreporting_oracle_specs ADD COLUMN is_bootstrap_peer bool NOT NULL;
        ALTER TABLE offchainreporting_oracle_specs ADD COLUMN encrypted_ocr_key_bundle_id bytea NOT NULL REFERENCES encrypted_ocr_key_bundles (id);
        ALTER TABLE offchainreporting_oracle_specs ADD COLUMN monitoring_endpoint TEXT;
        ALTER TABLE offchainreporting_oracle_specs ADD COLUMN transmitter_address bytea;  -- NOT NULL REFERENCES keys (address)
        ALTER TABLE offchainreporting_oracle_specs ADD COLUMN observation_timeout bigint NOT NULL;
        ALTER TABLE offchainreporting_oracle_specs ADD COLUMN blockchain_timeout bigint NOT NULL;
        ALTER TABLE offchainreporting_oracle_specs ADD COLUMN contract_config_tracker_subscribe_interval bigint NOT NULL;
        ALTER TABLE offchainreporting_oracle_specs ADD COLUMN contract_config_tracker_poll_interval bigint NOT NULL;
        ALTER TABLE offchainreporting_oracle_specs ADD COLUMN contract_config_confirmations INT NOT NULL;
        ALTER TABLE offchainreporting_oracle_specs ADD COLUMN created_at timestamptz NOT NULL;
        ALTER TABLE offchainreporting_oracle_specs ADD COLUMN updated_at timestamptz NOT NULL;
        ALTER TABLE offchainreporting_oracle_specs ADD CONSTRAINT chk_contract_address_length CHECK (octet_length(contract_address) = 20);

        -- NOTE: This will be extended with new IDs when we bring directrequest and fluxmonitor under the new jobspawner umbrella
        -- Only ONE id should ever be present
        CREATE TABLE jobs (
            id SERIAL PRIMARY KEY,
            pipeline_spec_id INT REFERENCES pipeline_specs (id),
            offchainreporting_oracle_spec_id INT REFERENCES offchainreporting_oracle_specs (id),
            CONSTRAINT chk_valid CHECK (
                offchainreporting_oracle_spec_id IS NOT NULL
            )
        );

        ---
        --- Notify the Chainlink node when a new job spec is created
        ---

        CREATE OR REPLACE FUNCTION notifyJobCreated() RETURNS TRIGGER AS $_$
        BEGIN
            PERFORM pg_notify('insert_on_jobs', NEW.id::text);
            RETURN NEW;
        END
        $_$ LANGUAGE 'plpgsql';

        CREATE TRIGGER notify_job_created
        AFTER INSERT ON jobs
        FOR EACH ROW EXECUTE PROCEDURE notifyJobCreated();



        CREATE UNIQUE INDEX idx_jobs_unique_offchainreporting_oracle_spec_ids ON jobs (offchainreporting_oracle_spec_id);

        CREATE UNIQUE INDEX idx_offchainreporting_oracle_specs_unique_key_bundles ON offchainreporting_oracle_specs (encrypted_ocr_key_bundle_id, contract_address);
        CREATE UNIQUE INDEX idx_offchainreporting_oracle_specs_unique_peer_ids ON offchainreporting_oracle_specs (p2p_peer_id, contract_address);
        -- CREATE UNIQUE INDEX idx_offchainreporting_oracle_specs_unique_job_ids ON offchainreporting_oracle_specs (job_id);
        -- CREATE INDEX idx_offchainreporting_oracle_specs_data_fetch_pipeline_spec_id ON offchainreporting_oracle_specs (data_fetch_pipeline_spec_id);

        CREATE INDEX idx_offchainreporting_oracle_specs_created_at ON offchainreporting_oracle_specs USING BRIN (created_at);
        CREATE INDEX idx_offchainreporting_oracle_specs_updated_at ON offchainreporting_oracle_specs USING BRIN (updated_at);


        ALTER TABLE log_consumptions ADD COLUMN job_id_v2 INT REFERENCES jobs (id) ON DELETE CASCADE;
        ALTER TABLE log_consumptions ALTER job_id DROP NOT NULL;
        ALTER TABLE log_consumptions ADD CONSTRAINT chk_log_consumptions_exactly_one_job_id CHECK (
            job_id IS NOT NULL AND job_id_v2 IS NULL
            OR
            job_id_v2 IS NOT NULL AND job_id IS NULL
        );
        DROP INDEX log_consumptions_unique_idx;
        CREATE UNIQUE INDEX log_consumptions_unique_idx ON log_consumptions ("job_id", "job_id_v2", "block_hash", "log_index");

    `).Error
}
