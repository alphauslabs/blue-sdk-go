package cost

// Deprecated compatibility aliases for renamed proto message types.
//
// AwsRispExportOptions was renamed to AwsExportReportOptions to reflect that
// the ExportReport API now supports all registered AWS export-report query
// domains.

// Migration:
//   - Replace AwsRispExportOptions with AwsExportReportOptions
//   - Replace AwsRispExportOptions_Columns with AwsExportReportOptions_Columns

// Deprecated: use AwsExportReportOptions instead.
type AwsRispExportOptions = AwsExportReportOptions

// Deprecated: use AwsExportReportOptions_Columns instead.
type AwsRispExportOptions_Columns = AwsExportReportOptions_Columns
