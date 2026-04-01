import { z } from "zod";

export const billtopConfigSchema = z.object({
	projectId: z.string().min(1),
	billingAccountId: z
		.string()
		.regex(/^[A-Z0-9]{6}-[A-Z0-9]{6}-[A-Z0-9]{6}$/, "Must be format XXXXXX-XXXXXX-XXXXXX"),
	datasetId: z.string().default("billtop_billing"),
	datasetLocation: z.enum(["US", "EU"]).default("US"),
});

export type BilltopConfig = z.infer<typeof billtopConfigSchema>;
