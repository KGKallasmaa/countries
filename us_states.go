package countries

// USStateCode represents the standard ANSI/FIPS two-letter state codes for US states and territories
type USStateCode string

const (
	USStateCodeAL USStateCode = "AL" // Alabama
	USStateCodeAK USStateCode = "AK" // Alaska
	USStateCodeAZ USStateCode = "AZ" // Arizona
	USStateCodeAR USStateCode = "AR" // Arkansas
	USStateCodeCA USStateCode = "CA" // California
	USStateCodeCO USStateCode = "CO" // Colorado
	USStateCodeCT USStateCode = "CT" // Connecticut
	USStateCodeDE USStateCode = "DE" // Delaware
	USStateCodeFL USStateCode = "FL" // Florida
	USStateCodeGA USStateCode = "GA" // Georgia
	USStateCodeHI USStateCode = "HI" // Hawaii
	USStateCodeID USStateCode = "ID" // Idaho
	USStateCodeIL USStateCode = "IL" // Illinois
	USStateCodeIN USStateCode = "IN" // Indiana
	USStateCodeIA USStateCode = "IA" // Iowa
	USStateCodeKS USStateCode = "KS" // Kansas
	USStateCodeKY USStateCode = "KY" // Kentucky
	USStateCodeLA USStateCode = "LA" // Louisiana
	USStateCodeME USStateCode = "ME" // Maine
	USStateCodeMD USStateCode = "MD" // Maryland
	USStateCodeMA USStateCode = "MA" // Massachusetts
	USStateCodeMI USStateCode = "MI" // Michigan
	USStateCodeMN USStateCode = "MN" // Minnesota
	USStateCodeMS USStateCode = "MS" // Mississippi
	USStateCodeMO USStateCode = "MO" // Missouri
	USStateCodeMT USStateCode = "MT" // Montana
	USStateCodeNE USStateCode = "NE" // Nebraska
	USStateCodeNV USStateCode = "NV" // Nevada
	USStateCodeNH USStateCode = "NH" // New Hampshire
	USStateCodeNJ USStateCode = "NJ" // New Jersey
	USStateCodeNM USStateCode = "NM" // New Mexico
	USStateCodeNY USStateCode = "NY" // New York
	USStateCodeNC USStateCode = "NC" // North Carolina
	USStateCodeND USStateCode = "ND" // North Dakota
	USStateCodeOH USStateCode = "OH" // Ohio
	USStateCodeOK USStateCode = "OK" // Oklahoma
	USStateCodeOR USStateCode = "OR" // Oregon
	USStateCodePA USStateCode = "PA" // Pennsylvania
	USStateCodeRI USStateCode = "RI" // Rhode Island
	USStateCodeSC USStateCode = "SC" // South Carolina
	USStateCodeSD USStateCode = "SD" // South Dakota
	USStateCodeTN USStateCode = "TN" // Tennessee
	USStateCodeTX USStateCode = "TX" // Texas
	USStateCodeUT USStateCode = "UT" // Utah
	USStateCodeVT USStateCode = "VT" // Vermont
	USStateCodeVA USStateCode = "VA" // Virginia
	USStateCodeWA USStateCode = "WA" // Washington
	USStateCodeWV USStateCode = "WV" // West Virginia
	USStateCodeWI USStateCode = "WI" // Wisconsin
	USStateCodeWY USStateCode = "WY" // Wyoming

	// US Territories and Federal District
	USStateCodeDC USStateCode = "DC" // District of Columbia
	USStateCodeAS USStateCode = "AS" // American Samoa
	USStateCodeGU USStateCode = "GU" // Guam
	USStateCodeMP USStateCode = "MP" // Northern Mariana Islands
	USStateCodePR USStateCode = "PR" // Puerto Rico
	USStateCodeVI USStateCode = "VI" // U.S. Virgin Islands
)

// AllUSStates contains all US state and territory codes
var (
	AllUSStates = []USStateCode{
		USStateCodeAK,
		USStateCodeAL,
		USStateCodeAR,
		USStateCodeAS,
		USStateCodeAZ,
		USStateCodeCA,
		USStateCodeCO,
		USStateCodeCT,
		USStateCodeDC,
		USStateCodeDE,
		USStateCodeFL,
		USStateCodeGA,
		USStateCodeGU,
		USStateCodeHI,
		USStateCodeIA,
		USStateCodeID,
		USStateCodeIL,
		USStateCodeIN,
		USStateCodeKS,
		USStateCodeKY,
		USStateCodeLA,
		USStateCodeMA,
		USStateCodeMD,
		USStateCodeME,
		USStateCodeMI,
		USStateCodeMN,
		USStateCodeMO,
		USStateCodeMP,
		USStateCodeMS,
		USStateCodeMT,
		USStateCodeNC,
		USStateCodeND,
		USStateCodeNE,
		USStateCodeNH,
		USStateCodeNJ,
		USStateCodeNM,
		USStateCodeNV,
		USStateCodeNY,
		USStateCodeOH,
		USStateCodeOK,
		USStateCodeOR,
		USStateCodePA,
		USStateCodePR,
		USStateCodeRI,
		USStateCodeSC,
		USStateCodeSD,
		USStateCodeTN,
		USStateCodeTX,
		USStateCodeUT,
		USStateCodeVA,
		USStateCodeVI,
		USStateCodeVT,
		USStateCodeWA,
		USStateCodeWI,
		USStateCodeWV,
		USStateCodeWY,
	}
)
