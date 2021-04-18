package util

var SVGLibrary = map[string]string{
	"app":     appIcon,
	"profile": profileIcon,
}

var appIcon = `<svg xmlns="http://www.w3.org/2000/svg" viewBox="-544 556 90 90" enable-background="new -549 551 100 100" style="width: 32px;height: 32px;"><g><path class="svg-fill" d="M-463.057,600.751l-0.186-1.845c-0.107-1.056,0.491-2.041,1.467-2.46c1.448-0.622,2.89-1.296,4.323-2.023 c0.994-0.504,1.506-1.643,1.253-2.728c-0.45-1.936-1.116-3.834-1.741-5.734c-0.354-1.076-1.412-1.781-2.539-1.662 c-1.579,0.167-3.141,0.383-4.68,0.644c-1.025,0.174-2.042-0.305-2.538-1.218l-1.025-1.888c-0.304-0.614-0.726-1.148-1.169-1.667 c-0.704-0.826-0.749-2.009-0.145-2.91c0.873-1.301,1.711-2.642,2.513-4.022c0.555-0.955,0.374-2.175-0.409-2.953 c-1.421-1.411-3.044-2.624-4.617-3.881c-0.871-0.696-2.115-0.721-2.987-0.026c-1.23,0.98-2.437,2.018-3.594,3.073 c-0.767,0.699-1.884,0.849-2.796,0.353l-1.888-1.025c-0.563-0.378-1.19-0.601-1.832-0.79c-1.059-0.312-1.766-1.282-1.745-2.386 c0.03-1.569,0.008-3.155-0.067-4.753c-0.05-1.081-0.832-2.011-1.887-2.251c-1.989-0.452-4.052-0.528-6.092-0.704 c-1.089-0.094-2.124,0.541-2.482,1.573c-0.507,1.46-0.962,2.967-1.363,4.504c-0.258,0.987-1.085,1.719-2.1,1.821l-2.179,0.22 c-0.726,0.042-1.446,0.118-2.153,0.26c-1.019,0.204-2.04-0.216-2.586-1.099c-0.813-1.317-1.679-2.619-2.595-3.904 c-0.629-0.882-1.785-1.224-2.808-0.866c-1.936,0.679-3.754,1.634-5.556,2.628c-0.948,0.523-1.449,1.615-1.198,2.668 c0.366,1.536,0.771,3.028,1.254,4.54c0.305,0.953,0.023,1.998-0.749,2.635l-1.715,1.414c-0.593,0.502-1.218,0.971-1.785,1.498 c-0.719,0.668-1.78,0.774-2.666,0.352c-1.391-0.664-2.818-1.288-4.277-1.871c-1.005-0.401-2.16-0.066-2.826,0.787 c-1.263,1.617-2.301,3.389-3.274,5.205c-0.509,0.951-0.343,2.136,0.435,2.884c1.146,1.102,2.311,2.151,3.521,3.161 c0.758,0.632,1.076,1.647,0.783,2.59l-0.67,2.152c-0.183,0.762-0.496,1.499-0.715,2.255c-0.268,0.925-1.065,1.577-2.02,1.699 c-1.528,0.196-3.065,0.443-4.606,0.745c-1.021,0.2-1.828,1.048-1.934,2.083c-0.217,2.116-0.217,4.229,0,6.343 c0.106,1.034,0.912,1.882,1.932,2.082c1.544,0.303,3.083,0.551,4.612,0.747c0.954,0.122,1.75,0.774,2.018,1.697 c0.22,0.756,0.533,1.493,0.716,2.258l0.668,2.151c0.293,0.943-0.027,1.957-0.785,2.589c-1.21,1.008-2.374,2.058-3.52,3.161 c-0.778,0.748-0.944,1.934-0.434,2.885c0.973,1.815,2.01,3.586,3.273,5.202c0.666,0.853,1.822,1.188,2.826,0.787 c1.46-0.583,2.888-1.208,4.279-1.872c0.886-0.423,1.947-0.318,2.666,0.35c0.567,0.527,1.193,0.996,1.786,1.498l1.718,1.418 c0.772,0.637,1.053,1.682,0.748,2.635c-0.485,1.512-0.89,3.004-1.257,4.54c-0.251,1.053,0.25,2.144,1.198,2.667 c1.803,0.995,3.622,1.949,5.56,2.628c1.022,0.358,2.179,0.016,2.807-0.867c0.916-1.285,1.781-2.588,2.595-3.904 c0.546-0.882,1.566-1.302,2.583-1.099c0.707,0.141,1.427,0.217,2.156,0.261l2.175,0.219c1.014,0.102,1.841,0.834,2.1,1.821 c0.403,1.538,0.858,3.045,1.367,4.506c0.359,1.031,1.394,1.666,2.482,1.572c2.039-0.176,4.101-0.252,6.09-0.703 c1.056-0.24,1.838-1.169,1.889-2.251c0.075-1.6,0.096-3.186,0.066-4.755c-0.021-1.103,0.683-2.072,1.741-2.385 c0.642-0.19,1.27-0.413,1.834-0.789l1.889-1.024c0.912-0.495,2.028-0.345,2.795,0.354c1.158,1.055,2.367,2.093,3.599,3.073 c0.872,0.693,2.115,0.667,2.985-0.028c1.573-1.257,3.194-2.471,4.615-3.884c0.782-0.778,0.962-1.998,0.408-2.952 c-0.801-1.379-1.64-2.72-2.513-4.02c-0.605-0.901-0.56-2.085,0.144-2.911c0.444-0.52,0.865-1.055,1.169-1.67l1.024-1.885 c0.496-0.914,1.514-1.392,2.539-1.217c1.539,0.262,3.101,0.477,4.68,0.643c1.125,0.118,2.182-0.585,2.536-1.66 c0.621-1.89,1.284-3.78,1.735-5.707c0.257-1.095-0.241-2.247-1.244-2.756c-1.433-0.727-2.875-1.402-4.324-2.024 c-0.976-0.419-1.574-1.403-1.468-2.459l0.186-1.856C-463.04,601.075-463.04,600.912-463.057,600.751z M-487.342,625.506 c-4.694,1.774-9.917,3.196-15.016,2.087c-5.007-0.798-10.017-2.703-13.757-6.29c-3.83-3.352-6.971-7.739-8.177-12.736 c-1.566-4.869-1.566-10.266,0-15.136c1.208-4.998,4.349-9.382,8.177-12.736c3.741-3.585,8.752-5.493,13.754-6.29 c5.101-1.112,10.321,0.312,15.018,2.087c4.739,2.07,8.533,5.79,11.492,9.938c2.798,4.29,3.953,9.305,4.325,14.362 c0.01,0.137,0.01,0.275,0,0.412c-0.369,5.058-1.527,10.074-4.325,14.362C-478.809,619.713-482.601,623.436-487.342,625.506z"/><g><path class="svg-fill" d="M-499,609.867c-0.706,0-1.28,0.573-1.28,1.28v4.49c0,0.706,0.573,1.28,1.28,1.28s1.28-0.573,1.28-1.28v-4.49C-497.72,610.44-498.294,609.867-499,609.867z"/><path class="svg-fill" d="M-485.251,594.416c0-6.023-3.857-11.292-9.598-13.108c-0.388-0.123-0.813-0.053-1.142,0.187c-0.329,0.241-0.523,0.624-0.523,1.032v12.057l-2.486,1.905l-2.486-1.905v-12.057c0-0.408-0.194-0.792-0.523-1.032c-0.33-0.241-0.752-0.311-1.143-0.187c-5.74,1.817-9.597,7.085-9.597,13.108c0,5.511,3.267,10.262,7.959,12.453v12.601c0,0.706,0.573,1.28,1.28,1.28c0.706,0,1.28-0.573,1.28-1.28v-11.704c1.038,0.251,2.116,0.399,3.231,0.399s2.193-0.148,3.232-0.399v11.705c0,0.706,0.573,1.28,1.28,1.28c0.706,0,1.28-0.573,1.28-1.28v-12.602C-488.517,604.677-485.251,599.927-485.251,594.416z M-501.196,605.395c-4.517-0.864-8.115-4.581-8.845-9.122c-0.808-5.024,1.727-9.706,5.996-11.85v10.791c0,0.398,0.185,0.774,0.502,1.016l3.766,2.886c0.459,0.352,1.097,0.352,1.556,0 l3.766-2.886c0.316-0.242,0.502-0.618,0.502-1.016v-10.791c3.736,1.876,6.145,5.698,6.145,9.992 C-487.81,601.304-494.067,606.759-501.196,605.395z"/></g></g></svg>`

var profileIcon = `<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20" style="height: 24px;width: 24px;"><circle class="svg-stroke" fill="none" stroke-width="1.1" cx="9.9" cy="6.4" r="4.4"></circle><path class="svg-stroke" fill="none" stroke-width="1.1" d="M1.5,19 C2.3,14.5 5.8,11.2 10,11.2 C14.2,11.2 17.7,14.6 18.5,19.2"></path></svg>`
