import { expect, test } from "bun:test";
import { version } from "./version.js";

test("version is a valid semver string", () => {
	expect(version).toMatch(/^\d+\.\d+\.\d+$/);
});
