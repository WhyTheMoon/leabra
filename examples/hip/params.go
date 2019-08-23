// File generated by params.SaveGoCode

package main

import "github.com/emer/emergent/params"

var SavedParamsSets = params.Sets{
	{Name: "Base", Desc: "these are the best params", Sheets: params.Sheets{
		"Network": &params.Sheet{
			{Sel: "Prjn", Desc: "keeping default params for generic prjns",
				Params: params.Params{
					"Prjn.Learn.Momentum.On": "true",
					"Prjn.Learn.Norm.On": "true",
					"Prjn.Learn.WtBal.On": "false",
				}},
			{Sel: ".EcCa1Prjn", Desc: "encoder projections -- no norm, moment",
				Params: params.Params{
					"Prjn.Learn.Lrate": "0.04",
					"Prjn.Learn.Momentum.On": "false",
					"Prjn.Learn.Norm.On": "false",
					"Prjn.Learn.WtBal.On": "false",
				}},
			{Sel: ".HippoCHL", Desc: "hippo CHL projections -- no norm, moment, but YES wtbal = sig better",
				Params: params.Params{
					"Prjn.CHL.Hebb": "0.05",
					"Prjn.Learn.Lrate": "0.4",
					"Prjn.Learn.Momentum.On": "false",
					"Prjn.Learn.Norm.On": "false",
					"Prjn.Learn.WtBal.On": "true",
				}},
			{Sel: "#CA1ToECout", Desc: "extra strong from CA1 to ECout",
				Params: params.Params{
					"Prjn.WtScale.Abs": "4.0",
				}},
			{Sel: "#InputToECin", Desc: "one-to-one input to EC",
				Params: params.Params{
					"Prjn.Learn.Learn": "false",
					"Prjn.WtInit.Mean": "0.8",
					"Prjn.WtInit.Var": "0.0",
				}},
			{Sel: "#ECoutToECin", Desc: "one-to-one out to in",
				Params: params.Params{
					"Prjn.Learn.Learn": "false",
					"Prjn.WtInit.Mean": "0.9",
					"Prjn.WtInit.Var": "0.01",
					"Prjn.WtScale.Rel": "0.5",
				}},
			{Sel: "#DGToCA3", Desc: "Mossy fibers: strong, non-learning",
				Params: params.Params{
					"Prjn.CHL.Hebb": "0.001",
					"Prjn.CHL.SAvgCor": "1",
					"Prjn.Learn.Learn": "false",
					"Prjn.WtInit.Mean": "0.9",
					"Prjn.WtInit.Var": "0.01",
					"Prjn.WtScale.Rel": "8",
				}},
			{Sel: "#CA3ToCA3", Desc: "CA3 recurrent cons",
				Params: params.Params{
					"Prjn.CHL.Hebb": "0.01",
					"Prjn.CHL.SAvgCor": "1",
					"Prjn.WtScale.Rel": "2",
				}},
			{Sel: "#CA3ToCA1", Desc: "Schaffer collaterals -- slower, less hebb",
				Params: params.Params{
					"Prjn.CHL.Hebb": "0.005",
					"Prjn.CHL.SAvgCor": "0.4",
					"Prjn.Learn.Lrate": "0.1",
				}},
			{Sel: ".EC", Desc: "all EC layers: only pools, no layer-level",
				Params: params.Params{
					"Layer.Act.Gbar.L": ".1",
					"Layer.Inhib.ActAvg.Init": "0.2",
					"Layer.Inhib.Layer.On": "false",
					"Layer.Inhib.Pool.Gi": "2.0",
					"Layer.Inhib.Pool.On": "true",
				}},
			{Sel: "#DG", Desc: "very sparse = high inibhition",
				Params: params.Params{
					"Layer.Inhib.ActAvg.Init": "0.01",
					"Layer.Inhib.Layer.Gi": "3.6",
				}},
			{Sel: "#CA3", Desc: "sparse = high inibhition",
				Params: params.Params{
					"Layer.Inhib.ActAvg.Init": "0.02",
					"Layer.Inhib.Layer.Gi": "2.8",
				}},
			{Sel: "#CA1", Desc: "CA1 only Pools",
				Params: params.Params{
					"Layer.Inhib.ActAvg.Init": "0.1",
					"Layer.Inhib.Layer.On": "false",
					"Layer.Inhib.Pool.Gi": "2.2",
					"Layer.Inhib.Pool.On": "true",
				}},
		},
		"Sim": &params.Sheet{
			{Sel: "Sim", Desc: "best params always finish in this time",
				Params: params.Params{
					"Sim.MaxEpcs": "10",
				}},
		},
	}},
	{Name: "NoCHL", Desc: "no learning in CHL main hip pathways -- for debugging auto-encoder", Sheets: params.Sheets{
		"Network": &params.Sheet{
			{Sel: ".HippoCHL", Desc: "no learning",
				Params: params.Params{
					"Prjn.Learn.Lrate": "0",
				}},
		},
		"Sim": &params.Sheet{
		},
	}},
	{Name: "CHLWtBal", Desc: "CHL uses weight balance -- much better -- now in base", Sheets: params.Sheets{
		"Network": &params.Sheet{
			{Sel: ".HippoCHL", Desc: "wtbal on",
				Params: params.Params{
					"Prjn.Learn.WtBal.On": "true",
				}},
		},
		"Sim": &params.Sheet{
		},
	}},
	{Name: "EncWtBal", Desc: "encoder uses weight balance -- worse", Sheets: params.Sheets{
		"Network": &params.Sheet{
			{Sel: ".EcCa1Prjn", Desc: "wtbal on",
				Params: params.Params{
					"Prjn.Learn.WtBal.On": "true",
				}},
		},
		"Sim": &params.Sheet{
		},
	}},
	{Name: "EncMom", Desc: "encoder uses momentum -- worse", Sheets: params.Sheets{
		"Network": &params.Sheet{
			{Sel: ".EcCa1Prjn", Desc: "moment on",
				Params: params.Params{
					"Prjn.Learn.Momentum.On": "true",
					"Prjn.Learn.Norm.On": "true",
				}},
		},
		"Sim": &params.Sheet{
		},
	}},
	{Name: "AllWtBal", Desc: "All use weight balance", Sheets: params.Sheets{
		"Network": &params.Sheet{
			{Sel: ".HippoCHL", Desc: "wtbal on",
				Params: params.Params{
					"Prjn.Learn.WtBal.On": "true",
				}},
			{Sel: ".EcCa1Prjn", Desc: "wtbal on",
				Params: params.Params{
					"Prjn.Learn.WtBal.On": "true",
				}},
		},
		"Sim": &params.Sheet{
		},
	}},
}
