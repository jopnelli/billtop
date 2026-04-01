/** Aggregated cost for a single GCP service. */
export interface CostByService {
	serviceName: string;
	grossCost: number;
	totalCredits: number;
	netCost: number;
	currency: string;
}

/** Aggregated cost for a single GCP project. */
export interface CostByProject {
	projectId: string;
	projectName: string;
	grossCost: number;
	netCost: number;
	currency: string;
}

/** Summary totals for a period. */
export interface CostSummary {
	grossCost: number;
	totalCredits: number;
	netCost: number;
	currency: string;
}

/** Time range for cost queries. */
export interface Period {
	start: Date;
	end: Date;
}
