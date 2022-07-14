# Temporal 

This is a template for running a production-ready Temporal cluster on Render. The setup supports independent autoscaling for each Temporal service (frontend, matching, history, worker), has [advanced visibility](https://docs.temporal.io/docs/content/what-is-advanced-visibility/) backed by Elasticsearch, and includes an example Go app to trigger and run workflows. Create a new repo using this template, and then click the button below to try it out:

[![Deploy to Render](https://render.com/images/deploy-to-render-button.svg)](https://render.com/deploy?repo=https://github.com/render-examples/temporal)

For deploy instructions, see our [Temporal guide](https://render.com/docs/deploy-temporal).

# Acknowledgements

[auto-setup-override.sh](temporal-cluster/server/auto-setup/auto-setup-override.sh) is based on Temporal's [auto-setup.sh script](https://github.com/temporalio/temporal/blob/077d39c775/docker/auto-setup.sh), with some modifications made to better accommodate Render's architecture.
