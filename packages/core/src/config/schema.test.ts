import { describe, expect, test } from "bun:test";
import { billtopConfigSchema } from "./schema.js";

describe("billtopConfigSchema", () => {
	test("accepts valid config", () => {
		const result = billtopConfigSchema.parse({
			projectId: "my-project",
			billingAccountId: "AABBCC-DDEEFF-001122",
		});
		expect(result.projectId).toBe("my-project");
		expect(result.datasetId).toBe("billtop_billing");
		expect(result.datasetLocation).toBe("US");
	});

	test("rejects empty projectId", () => {
		expect(() =>
			billtopConfigSchema.parse({
				projectId: "",
				billingAccountId: "AABBCC-DDEEFF-001122",
			}),
		).toThrow();
	});

	test("rejects invalid billing account format", () => {
		expect(() =>
			billtopConfigSchema.parse({
				projectId: "my-project",
				billingAccountId: "not-a-billing-id",
			}),
		).toThrow();
	});

	test("accepts custom dataset settings", () => {
		const result = billtopConfigSchema.parse({
			projectId: "my-project",
			billingAccountId: "AABBCC-DDEEFF-001122",
			datasetId: "custom_billing",
			datasetLocation: "EU",
		});
		expect(result.datasetId).toBe("custom_billing");
		expect(result.datasetLocation).toBe("EU");
	});

	test("rejects unknown dataset location", () => {
		expect(() =>
			billtopConfigSchema.parse({
				projectId: "my-project",
				billingAccountId: "AABBCC-DDEEFF-001122",
				datasetLocation: "asia",
			}),
		).toThrow();
	});
});
