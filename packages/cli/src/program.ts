import { version } from "@billtop/core";
import { Command } from "commander";

export function createProgram(): Command {
	const program = new Command();

	program
		.name("billtop")
		.description("Open source GCP cost optimization")
		.version(version, "-v, --version");

	return program;
}
