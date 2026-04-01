import { describe, expect, test } from "bun:test";
import { createProgram } from "./program.js";

describe("createProgram", () => {
	test("has correct name", () => {
		const program = createProgram();
		expect(program.name()).toBe("billtop");
	});

	test("outputs version with --version flag", () => {
		const program = createProgram();
		program.exitOverride();

		let output = "";
		program.configureOutput({
			writeOut: (str) => {
				output = str;
			},
		});

		try {
			program.parse(["--version"], { from: "user" });
		} catch (e: unknown) {
			if (!(e instanceof Error && e.constructor.name === "CommanderError")) throw e;
		}

		expect(output).toMatch(/^\d+\.\d+\.\d+/);
	});
});
