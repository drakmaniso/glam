// Copyright (c) 2013-2016 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package math

import (
	"math"
	"testing"
)

//------------------------------------------------------------------------------

func TestTan(t *testing.T) {
	var x float32
	for _, tt := range tan_tests {
		x = Tan(tt.in)
		if !IsNearlyEqual(x, tt.out, 2*EpsilonFloat32) {
			t.Errorf("Relative error for Tan(%.100g): %.100g instead of %.100g\n", tt.in, x, tt.out)
		}
		if !IsAlmostEqual(x, tt.out, 3) {
			t.Errorf("ULP error for Tan(%.100g): %.100g instead of %.100g\n", tt.in, x, tt.out)
		}
	}
}

//------------------------------------------------------------------------------

func BenchmarkTan_math64(b *testing.B) {
	a := float64(0.5)
	for i := 0; i < b.N; i++ {
		_ = math.Tan(a)
	}
}

//------------------------------------------------------------------------------

func BenchmarkTan_math32(b *testing.B) {
	a := float32(0.5)
	for i := 0; i < b.N; i++ {
		_ = float32(math.Tan(float64(a)))
	}
}

//------------------------------------------------------------------------------

func BenchmarkTan_glam(b *testing.B) {
	a := float32(0.5)
	for i := 0; i < b.N; i++ {
		_ = Tan(a)
	}
}

//------------------------------------------------------------------------------

var tan_tests = [...]struct {
	in  float32
	out float32
}{
	// Special cases

	{-Pi / 3, -1.7320508075688772935274463415058723669428052538103806280558069794519330169088000370811461867572485757},
	{-Pi / 4, -1},
	{-Pi / 6, -0.57735026918962576450914878050195745564760175127012687601860232648397767230293334569371539558574952523},
	{0, 0},
	{Pi / 6, 0.57735026918962576450914878050195745564760175127012687601860232648397767230293334569371539558574952523},
	{Pi / 4, 1},
	{Pi / 3, 1.7320508075688772935274463415058723669428052538103806280558069794519330169088000370811461867572485757},

	// The following values have been generated with Gnu MPFR, 24 bits mantissa, roundTiesToEven

	{-3.14000010e+00, +1.59255008e-03},
	{-3.13000011e+00, +1.15930587e-02},
	{-3.12000012e+00, +2.15958860e-02},
	{-3.11000013e+00, +3.16030346e-02},
	{-3.10000014e+00, +4.16165106e-02},
	{-3.09000015e+00, +5.16383275e-02},
	{-3.08000016e+00, +6.16704971e-02},
	{-3.07000017e+00, +7.17150494e-02},
	{-3.06000018e+00, +8.17740187e-02},
	{-3.05000019e+00, +9.18494537e-02},
	{-3.04000020e+00, +1.01943418e-01},
	{-3.03000021e+00, +1.12057976e-01},
	{-3.02000022e+00, +1.22195236e-01},
	{-3.01000023e+00, +1.32357299e-01},
	{-3.00000024e+00, +1.42546296e-01},
	{-2.99000025e+00, +1.52764395e-01},
	{-2.98000026e+00, +1.63013756e-01},
	{-2.97000027e+00, +1.73296586e-01},
	{-2.96000028e+00, +1.83615118e-01},
	{-2.95000029e+00, +1.93971604e-01},
	{-2.94000030e+00, +2.04368368e-01},
	{-2.93000031e+00, +2.14807704e-01},
	{-2.92000031e+00, +2.25291982e-01},
	{-2.91000032e+00, +2.35823616e-01},
	{-2.90000033e+00, +2.46405035e-01},
	{-2.89000034e+00, +2.57038742e-01},
	{-2.88000035e+00, +2.67727256e-01},
	{-2.87000036e+00, +2.78473139e-01},
	{-2.86000037e+00, +2.89279073e-01},
	{-2.85000038e+00, +3.00147682e-01},
	{-2.84000039e+00, +3.11081737e-01},
	{-2.83000040e+00, +3.22084039e-01},
	{-2.82000041e+00, +3.33157450e-01},
	{-2.81000042e+00, +3.44304889e-01},
	{-2.80000043e+00, +3.55529338e-01},
	{-2.79000044e+00, +3.66833925e-01},
	{-2.78000045e+00, +3.78221720e-01},
	{-2.77000046e+00, +3.89696032e-01},
	{-2.76000047e+00, +4.01260078e-01},
	{-2.75000048e+00, +4.12917346e-01},
	{-2.74000049e+00, +4.24671263e-01},
	{-2.73000050e+00, +4.36525434e-01},
	{-2.72000051e+00, +4.48483586e-01},
	{-2.71000051e+00, +4.60549444e-01},
	{-2.70000052e+00, +4.72727001e-01},
	{-2.69000053e+00, +4.85020220e-01},
	{-2.68000054e+00, +4.97433275e-01},
	{-2.67000055e+00, +5.09970427e-01},
	{-2.66000056e+00, +5.22636116e-01},
	{-2.65000057e+00, +5.35434902e-01},
	{-2.64000058e+00, +5.48371494e-01},
	{-2.63000059e+00, +5.61450779e-01},
	{-2.62000060e+00, +5.74677706e-01},
	{-2.61000061e+00, +5.88057578e-01},
	{-2.60000062e+00, +6.01595759e-01},
	{-2.59000063e+00, +6.15297794e-01},
	{-2.58000064e+00, +6.29169524e-01},
	{-2.57000065e+00, +6.43216908e-01},
	{-2.56000066e+00, +6.57446206e-01},
	{-2.55000067e+00, +6.71863794e-01},
	{-2.54000068e+00, +6.86476469e-01},
	{-2.53000069e+00, +7.01291144e-01},
	{-2.52000070e+00, +7.16315091e-01},
	{-2.51000071e+00, +7.31555820e-01},
	{-2.50000072e+00, +7.47021198e-01},
	{-2.49000072e+00, +7.62719393e-01},
	{-2.48000073e+00, +7.78658867e-01},
	{-2.47000074e+00, +7.94848561e-01},
	{-2.46000075e+00, +8.11297655e-01},
	{-2.45000076e+00, +8.28015864e-01},
	{-2.44000077e+00, +8.45013261e-01},
	{-2.43000078e+00, +8.62300396e-01},
	{-2.42000079e+00, +8.79888237e-01},
	{-2.41000080e+00, +8.97788346e-01},
	{-2.40000081e+00, +9.16012824e-01},
	{-2.39000082e+00, +9.34574246e-01},
	{-2.38000083e+00, +9.53485847e-01},
	{-2.37000084e+00, +9.72761631e-01},
	{-2.36000085e+00, +9.92416084e-01},
	{-2.35000086e+00, +1.01246464e+00},
	{-2.34000087e+00, +1.03292322e+00},
	{-2.33000088e+00, +1.05380893e+00},
	{-2.32000089e+00, +1.07513964e+00},
	{-2.31000090e+00, +1.09693384e+00},
	{-2.30000091e+00, +1.11921155e+00},
	{-2.29000092e+00, +1.14199364e+00},
	{-2.28000093e+00, +1.16530204e+00},
	{-2.27000093e+00, +1.18916023e+00},
	{-2.26000094e+00, +1.21359253e+00},
	{-2.25000095e+00, +1.23862517e+00},
	{-2.24000096e+00, +1.26428580e+00},
	{-2.23000097e+00, +1.29060352e+00},
	{-2.22000098e+00, +1.31760955e+00},
	{-2.21000099e+00, +1.34533679e+00},
	{-2.20000100e+00, +1.37382019e+00},
	{-2.19000101e+00, +1.40309715e+00},
	{-2.18000102e+00, +1.43320739e+00},
	{-2.17000103e+00, +1.46419334e+00},
	{-2.16000104e+00, +1.49610019e+00},
	{-2.15000105e+00, +1.52897620e+00},
	{-2.14000106e+00, +1.56287336e+00},
	{-2.13000107e+00, +1.59784675e+00},
	{-2.12000108e+00, +1.63395607e+00},
	{-2.11000109e+00, +1.67126501e+00},
	{-2.10000110e+00, +1.70984221e+00},
	{-2.09000111e+00, +1.74976170e+00},
	{-2.08000112e+00, +1.79110312e+00},
	{-2.07000113e+00, +1.83395243e+00},
	{-2.06000113e+00, +1.87840295e+00},
	{-2.05000114e+00, +1.92455530e+00},
	{-2.04000115e+00, +1.97251904e+00},
	{-2.03000116e+00, +2.02241325e+00},
	{-2.02000117e+00, +2.07436705e+00},
	{-2.01000118e+00, +2.12852216e+00},
	{-2.00000119e+00, +2.18503308e+00},
	{-1.99000120e+00, +2.24406862e+00},
	{-1.98000121e+00, +2.30581450e+00},
	{-1.97000122e+00, +2.37047553e+00},
	{-1.96000123e+00, +2.43827629e+00},
	{-1.95000124e+00, +2.50946641e+00},
	{-1.94000125e+00, +2.58432150e+00},
	{-1.93000126e+00, +2.66314840e+00},
	{-1.92000127e+00, +2.74628878e+00},
	{-1.91000128e+00, +2.83412504e+00},
	{-1.90000129e+00, +2.92708516e+00},
	{-1.89000130e+00, +3.02565169e+00},
	{-1.88000131e+00, +3.13036919e+00},
	{-1.87000132e+00, +3.24185491e+00},
	{-1.86000133e+00, +3.36081147e+00},
	{-1.85000134e+00, +3.48804212e+00},
	{-1.84000134e+00, +3.62446952e+00},
	{-1.83000135e+00, +3.77115870e+00},
	{-1.82000136e+00, +3.92934561e+00},
	{-1.81000137e+00, +4.10047293e+00},
	{-1.80000138e+00, +4.28623486e+00},
	{-1.79000139e+00, +4.48863506e+00},
	{-1.78000140e+00, +4.71005964e+00},
	{-1.77000141e+00, +4.95337439e+00},
	{-1.76000142e+00, +5.22205067e+00},
	{-1.75000143e+00, +5.52033472e+00},
	{-1.74000144e+00, +5.85347700e+00},
	{-1.73000145e+00, +6.22804642e+00},
	{-1.72000146e+00, +6.65237331e+00},
	{-1.71000147e+00, +7.13718081e+00},
	{-1.70000148e+00, +7.69651318e+00},
	{-1.69000149e+00, +8.34912586e+00},
	{-1.68000150e+00, +9.12064457e+00},
	{-1.67000151e+00, +1.00470285e+01},
	{-1.66000152e+00, +1.11803589e+01},
	{-1.65000153e+00, +1.25990210e+01},
	{-1.64000154e+00, +1.44267035e+01},
	{-1.63000154e+00, +1.68706627e+01},
	{-1.62000155e+00, +2.03066387e+01},
	{-1.61000156e+00, +2.54937267e+01},
	{-1.60000157e+00, +3.42306862e+01},
	{-1.59000158e+00, +5.20626755e+01},
	{-1.58000159e+00, +1.08630402e+02},
	{-1.57000160e+00, -1.25829724e+03},
	{-1.56000161e+00, -9.26343231e+01},
	{-1.55000162e+00, -4.80822334e+01},
	{-1.54000163e+00, -3.24628601e+01},
	{-1.53000164e+00, -2.44993973e+01},
	{-1.52000165e+00, -1.96701679e+01},
	{-1.51000166e+00, -1.64285412e+01},
	{-1.50000167e+00, -1.41017532e+01},
	{-1.49000168e+00, -1.23501139e+01},
	{-1.48000169e+00, -1.09835844e+01},
	{-1.47000170e+00, -9.88754272e+00},
	{-1.46000171e+00, -8.98874760e+00},
	{-1.45000172e+00, -8.23821068e+00},
	{-1.44000173e+00, -7.60192776e+00},
	{-1.43000174e+00, -7.05555201e+00},
	{-1.42000175e+00, -6.58119678e+00},
	{-1.41000175e+00, -6.16542482e+00},
	{-1.40000176e+00, -5.79794502e+00},
	{-1.39000177e+00, -5.47074366e+00},
	{-1.38000178e+00, -5.17748690e+00},
	{-1.37000179e+00, -4.91310310e+00},
	{-1.36000180e+00, -4.67348242e+00},
	{-1.35000181e+00, -4.45525932e+00},
	{-1.34000182e+00, -4.25565290e+00},
	{-1.33000183e+00, -4.07234192e+00},
	{-1.32000184e+00, -3.90337777e+00},
	{-1.31000185e+00, -3.74710870e+00},
	{-1.30000186e+00, -3.60212851e+00},
	{-1.29000187e+00, -3.46723008e+00},
	{-1.28000188e+00, -3.34137273e+00},
	{-1.27000189e+00, -3.22365475e+00},
	{-1.26000190e+00, -3.11328936e+00},
	{-1.25000191e+00, -3.00958896e+00},
	{-1.24000192e+00, -2.91194797e+00},
	{-1.23000193e+00, -2.81983304e+00},
	{-1.22000194e+00, -2.73277068e+00},
	{-1.21000195e+00, -2.65034032e+00},
	{-1.20000196e+00, -2.57216644e+00},
	{-1.19000196e+00, -2.49791360e+00},
	{-1.18000197e+00, -2.42727995e+00},
	{-1.17000198e+00, -2.35999417e+00},
	{-1.16000199e+00, -2.29581094e+00},
	{-1.15000200e+00, -2.23450899e+00},
	{-1.14000201e+00, -2.17588663e+00},
	{-1.13000202e+00, -2.11976123e+00},
	{-1.12000203e+00, -2.06596589e+00},
	{-1.11000204e+00, -2.01434851e+00},
	{-1.10000205e+00, -1.96476960e+00},
	{-1.09000206e+00, -1.91710150e+00},
	{-1.08000207e+00, -1.87122667e+00},
	{-1.07000208e+00, -1.82703722e+00},
	{-1.06000209e+00, -1.78443360e+00},
	{-1.05000210e+00, -1.74332380e+00},
	{-1.04000211e+00, -1.70362282e+00},
	{-1.03000212e+00, -1.66525197e+00},
	{-1.02000213e+00, -1.62813818e+00},
	{-1.01000214e+00, -1.59221363e+00},
	{-1.00000215e+00, -1.55741513e+00},
	{-9.90002155e-01, -1.52368391e+00},
	{-9.80002165e-01, -1.49096525e+00},
	{-9.70002174e-01, -1.45920789e+00},
	{-9.60002184e-01, -1.42836416e+00},
	{-9.50002193e-01, -1.39838910e+00},
	{-9.40002203e-01, -1.36924076e+00},
	{-9.30002213e-01, -1.34088004e+00},
	{-9.20002222e-01, -1.31326973e+00},
	{-9.10002232e-01, -1.28637528e+00},
	{-9.00002241e-01, -1.26016402e+00},
	{-8.90002251e-01, -1.23460519e+00},
	{-8.80002260e-01, -1.20966971e+00},
	{-8.70002270e-01, -1.18533027e+00},
	{-8.60002279e-01, -1.16156125e+00},
	{-8.50002289e-01, -1.13833797e+00},
	{-8.40002298e-01, -1.11563754e+00},
	{-8.30002308e-01, -1.09343803e+00},
	{-8.20002317e-01, -1.07171869e+00},
	{-8.10002327e-01, -1.05045998e+00},
	{-8.00002337e-01, -1.02964342e+00},
	{-7.90002346e-01, -1.00925100e+00},
	{-7.80002356e-01, -9.89266217e-01},
	{-7.70002365e-01, -9.69672918e-01},
	{-7.60002375e-01, -9.50455964e-01},
	{-7.50002384e-01, -9.31600928e-01},
	{-7.40002394e-01, -9.13093925e-01},
	{-7.30002403e-01, -8.94921839e-01},
	{-7.20002413e-01, -8.77072155e-01},
	{-7.10002422e-01, -8.59532893e-01},
	{-7.00002432e-01, -8.42292547e-01},
	{-6.90002441e-01, -8.25340211e-01},
	{-6.80002451e-01, -8.08665454e-01},
	{-6.70002460e-01, -7.92258203e-01},
	{-6.60002470e-01, -7.76108861e-01},
	{-6.50002480e-01, -7.60208309e-01},
	{-6.40002489e-01, -7.44547665e-01},
	{-6.30002499e-01, -7.29118586e-01},
	{-6.20002508e-01, -7.13912785e-01},
	{-6.10002518e-01, -6.98922634e-01},
	{-6.00002527e-01, -6.84140503e-01},
	{-5.90002537e-01, -6.69559300e-01},
	{-5.80002546e-01, -6.55172110e-01},
	{-5.70002556e-01, -6.40972137e-01},
	{-5.60002565e-01, -6.26953125e-01},
	{-5.50002575e-01, -6.13108754e-01},
	{-5.40002584e-01, -5.99433124e-01},
	{-5.30002594e-01, -5.85920513e-01},
	{-5.20002604e-01, -5.72565258e-01},
	{-5.10002613e-01, -5.59362173e-01},
	{-5.00002623e-01, -5.46305895e-01},
	{-4.90002632e-01, -5.33391535e-01},
	{-4.80002642e-01, -5.20614207e-01},
	{-4.70002651e-01, -5.07969260e-01},
	{-4.60002661e-01, -4.95452076e-01},
	{-4.50002670e-01, -4.83058363e-01},
	{-4.40002680e-01, -4.70783800e-01},
	{-4.30002689e-01, -4.58624274e-01},
	{-4.20002699e-01, -4.46575791e-01},
	{-4.10002708e-01, -4.34634417e-01},
	{-4.00002718e-01, -4.22796428e-01},
	{-3.90002728e-01, -4.11058098e-01},
	{-3.80002737e-01, -3.99415880e-01},
	{-3.70002747e-01, -3.87866318e-01},
	{-3.60002756e-01, -3.76405984e-01},
	{-3.50002766e-01, -3.65031630e-01},
	{-3.40002775e-01, -3.53740007e-01},
	{-3.30002785e-01, -3.42527986e-01},
	{-3.20002794e-01, -3.31392497e-01},
	{-3.10002804e-01, -3.20330590e-01},
	{-3.00002813e-01, -3.09339345e-01},
	{-2.90002823e-01, -2.98415869e-01},
	{-2.80002832e-01, -2.87557393e-01},
	{-2.70002842e-01, -2.76761204e-01},
	{-2.60002851e-01, -2.66024590e-01},
	{-2.50002861e-01, -2.55344957e-01},
	{-2.40002856e-01, -2.44719729e-01},
	{-2.30002850e-01, -2.34146371e-01},
	{-2.20002845e-01, -2.23622411e-01},
	{-2.10002840e-01, -2.13145420e-01},
	{-2.00002834e-01, -2.02712983e-01},
	{-1.90002829e-01, -1.92322776e-01},
	{-1.80002823e-01, -1.81972444e-01},
	{-1.70002818e-01, -1.71659723e-01},
	{-1.60002813e-01, -1.61382347e-01},
	{-1.50002807e-01, -1.51138082e-01},
	{-1.40002802e-01, -1.40924752e-01},
	{-1.30002797e-01, -1.30740166e-01},
	{-1.20002799e-01, -1.20582178e-01},
	{-1.10002801e-01, -1.10448658e-01},
	{-1.00002803e-01, -1.00337505e-01},
	{-9.00028050e-02, -9.02466178e-02},
	{-8.00028071e-02, -8.01739320e-02},
	{-7.00028092e-02, -7.01173842e-02},
	{-6.00028113e-02, -6.00749254e-02},
	{-5.00028133e-02, -5.00445291e-02},
	{-4.00028154e-02, -4.00241688e-02},
	{-3.00028156e-02, -3.00118215e-02},
	{-2.00028159e-02, -2.00054832e-02},
	{-1.00028161e-02, -1.00031495e-02},
	{-2.81631947e-06, -2.81631947e-06},
	{+9.99718346e-03, +9.99751687e-03},
	{+1.99971832e-02, +1.99998487e-02},
	{+2.99971830e-02, +3.00061833e-02},
	{+3.99971828e-02, +4.00185250e-02},
	{+4.99971807e-02, +5.00388816e-02},
	{+5.99971786e-02, +6.00692704e-02},
	{+6.99971765e-02, +7.01117218e-02},
	{+7.99971744e-02, +8.01682621e-02},
	{+8.99971724e-02, +9.02409405e-02},
	{+9.99971703e-02, +1.00331813e-01},
	{+1.09997168e-01, +1.10442959e-01},
	{+1.19997166e-01, +1.20576464e-01},
	{+1.29997164e-01, +1.30734429e-01},
	{+1.39997169e-01, +1.40919015e-01},
	{+1.49997175e-01, +1.51132330e-01},
	{+1.59997180e-01, +1.61376566e-01},
	{+1.69997185e-01, +1.71653926e-01},
	{+1.79997191e-01, +1.81966633e-01},
	{+1.89997196e-01, +1.92316934e-01},
	{+1.99997202e-01, +2.02707127e-01},
	{+2.09997207e-01, +2.13139519e-01},
	{+2.19997212e-01, +2.23616496e-01},
	{+2.29997218e-01, +2.34140426e-01},
	{+2.39997223e-01, +2.44713753e-01},
	{+2.49997228e-01, +2.55338967e-01},
	{+2.59997219e-01, +2.66018569e-01},
	{+2.69997209e-01, +2.76755124e-01},
	{+2.79997200e-01, +2.87551284e-01},
	{+2.89997190e-01, +2.98409730e-01},
	{+2.99997181e-01, +3.09333175e-01},
	{+3.09997171e-01, +3.20324391e-01},
	{+3.19997162e-01, +3.31386268e-01},
	{+3.29997152e-01, +3.42521697e-01},
	{+3.39997143e-01, +3.53733659e-01},
	{+3.49997133e-01, +3.65025252e-01},
	{+3.59997123e-01, +3.76399577e-01},
	{+3.69997114e-01, +3.87859851e-01},
	{+3.79997104e-01, +3.99409354e-01},
	{+3.89997095e-01, +4.11051512e-01},
	{+3.99997085e-01, +4.22789782e-01},
	{+4.09997076e-01, +4.34627742e-01},
	{+4.19997066e-01, +4.46569026e-01},
	{+4.29997057e-01, +4.58617449e-01},
	{+4.39997047e-01, +4.70776916e-01},
	{+4.49997038e-01, +4.83051419e-01},
	{+4.59997028e-01, +4.95445073e-01},
	{+4.69997019e-01, +5.07962167e-01},
	{+4.79997009e-01, +5.20607054e-01},
	{+4.89997000e-01, +5.33384264e-01},
	{+4.99996990e-01, +5.46298563e-01},
	{+5.09997010e-01, +5.59354782e-01},
	{+5.19997001e-01, +5.72557867e-01},
	{+5.29996991e-01, +5.85912943e-01},
	{+5.39996982e-01, +5.99425495e-01},
	{+5.49996972e-01, +6.13101065e-01},
	{+5.59996963e-01, +6.26945317e-01},
	{+5.69996953e-01, +6.40964270e-01},
	{+5.79996943e-01, +6.55164063e-01},
	{+5.89996934e-01, +6.69551194e-01},
	{+5.99996924e-01, +6.84132278e-01},
	{+6.09996915e-01, +6.98914289e-01},
	{+6.19996905e-01, +7.13904321e-01},
	{+6.29996896e-01, +7.29110003e-01},
	{+6.39996886e-01, +7.44538963e-01},
	{+6.49996877e-01, +7.60199487e-01},
	{+6.59996867e-01, +7.76099920e-01},
	{+6.69996858e-01, +7.92249084e-01},
	{+6.79996848e-01, +8.08656156e-01},
	{+6.89996839e-01, +8.25330794e-01},
	{+6.99996829e-01, +8.42282951e-01},
	{+7.09996819e-01, +8.59523118e-01},
	{+7.19996810e-01, +8.77062261e-01},
	{+7.29996800e-01, +8.94911766e-01},
	{+7.39996791e-01, +9.13083673e-01},
	{+7.49996781e-01, +9.31590438e-01},
	{+7.59996772e-01, +9.50445294e-01},
	{+7.69996762e-01, +9.69662070e-01},
	{+7.79996753e-01, +9.89255130e-01},
	{+7.89996743e-01, +1.00923967e+00},
	{+7.99996734e-01, +1.02963185e+00},
	{+8.09996724e-01, +1.05044830e+00},
	{+8.19996715e-01, +1.07170665e+00},
	{+8.29996705e-01, +1.09342563e+00},
	{+8.39996696e-01, +1.11562490e+00},
	{+8.49996686e-01, +1.13832510e+00},
	{+8.59996676e-01, +1.16154802e+00},
	{+8.69996667e-01, +1.18531680e+00},
	{+8.79996657e-01, +1.20965588e+00},
	{+8.89996648e-01, +1.23459101e+00},
	{+8.99996638e-01, +1.26014948e+00},
	{+9.09996629e-01, +1.28636038e+00},
	{+9.19996619e-01, +1.31325448e+00},
	{+9.29996610e-01, +1.34086430e+00},
	{+9.39996600e-01, +1.36922467e+00},
	{+9.49996591e-01, +1.39837253e+00},
	{+9.59996581e-01, +1.42834711e+00},
	{+9.69996572e-01, +1.45919037e+00},
	{+9.79996562e-01, +1.49094713e+00},
	{+9.89996552e-01, +1.52366531e+00},
	{+9.99996543e-01, +1.55739594e+00},
	{+1.00999653e+00, +1.59219372e+00},
	{+1.01999652e+00, +1.62811768e+00},
	{+1.02999651e+00, +1.66523087e+00},
	{+1.03999650e+00, +1.70360100e+00},
	{+1.04999650e+00, +1.74330115e+00},
	{+1.05999649e+00, +1.78441012e+00},
	{+1.06999648e+00, +1.82701290e+00},
	{+1.07999647e+00, +1.87120140e+00},
	{+1.08999646e+00, +1.91707528e+00},
	{+1.09999645e+00, +1.96474242e+00},
	{+1.10999644e+00, +2.01432014e+00},
	{+1.11999643e+00, +2.06593657e+00},
	{+1.12999642e+00, +2.11973047e+00},
	{+1.13999641e+00, +2.17585444e+00},
	{+1.14999640e+00, +2.23447537e+00},
	{+1.15999639e+00, +2.29577589e+00},
	{+1.16999638e+00, +2.35995722e+00},
	{+1.17999637e+00, +2.42724133e+00},
	{+1.18999636e+00, +2.49787307e+00},
	{+1.19999635e+00, +2.57212377e+00},
	{+1.20999634e+00, +2.65029526e+00},
	{+1.21999633e+00, +2.73272324e+00},
	{+1.22999632e+00, +2.81978273e+00},
	{+1.23999631e+00, +2.91189504e+00},
	{+1.24999630e+00, +3.00953245e+00},
	{+1.25999629e+00, +3.11322951e+00},
	{+1.26999629e+00, +3.22359085e+00},
	{+1.27999628e+00, +3.34130478e+00},
	{+1.28999627e+00, +3.46715713e+00},
	{+1.29999626e+00, +3.60205007e+00},
	{+1.30999625e+00, +3.74702454e+00},
	{+1.31999624e+00, +3.90328670e+00},
	{+1.32999623e+00, +4.07224369e+00},
	{+1.33999622e+00, +4.25554562e+00},
	{+1.34999621e+00, +4.45514250e+00},
	{+1.35999620e+00, +4.67335463e+00},
	{+1.36999619e+00, +4.91296244e+00},
	{+1.37999618e+00, +5.17733097e+00},
	{+1.38999617e+00, +5.47057009e+00},
	{+1.39999616e+00, +5.79775095e+00},
	{+1.40999615e+00, +6.16520596e+00},
	{+1.41999614e+00, +6.58094835e+00},
	{+1.42999613e+00, +7.05526733e+00},
	{+1.43999612e+00, +7.60159826e+00},
	{+1.44999611e+00, +8.23782539e+00},
	{+1.45999610e+00, +8.98828888e+00},
	{+1.46999609e+00, +9.88698959e+00},
	{+1.47999609e+00, +1.09829035e+01},
	{+1.48999608e+00, +1.23492537e+01},
	{+1.49999607e+00, +1.41006336e+01},
	{+1.50999606e+00, +1.64270229e+01},
	{+1.51999605e+00, +1.96679955e+01},
	{+1.52999604e+00, +2.44960289e+01},
	{+1.53999603e+00, +3.24569511e+01},
	{+1.54999602e+00, +4.80692749e+01},
	{+1.55999601e+00, +9.25862656e+01},
	{+1.56999600e+00, +1.24948828e+03},
	{+1.57999599e+00, -1.08696564e+02},
	{+1.58999598e+00, -5.20778732e+01},
	{+1.59999597e+00, -3.42372589e+01},
	{+1.60999596e+00, -2.54973736e+01},
	{+1.61999595e+00, -2.03089561e+01},
	{+1.62999594e+00, -1.68722630e+01},
	{+1.63999593e+00, -1.44278755e+01},
	{+1.64999592e+00, -1.25999165e+01},
	{+1.65999591e+00, -1.11810656e+01},
	{+1.66999590e+00, -1.00475998e+01},
	{+1.67999589e+00, -9.12111664e+00},
	{+1.68999588e+00, -8.34952164e+00},
	{+1.69999588e+00, -7.69685078e+00},
	{+1.70999587e+00, -7.13747215e+00},
	{+1.71999586e+00, -6.65262699e+00},
	{+1.72999585e+00, -6.22826958e+00},
	{+1.73999584e+00, -5.85367489e+00},
	{+1.74999583e+00, -5.52051115e+00},
	{+1.75999582e+00, -5.22220898e+00},
	{+1.76999581e+00, -4.95351744e+00},
	{+1.77999580e+00, -4.71018934e+00},
	{+1.78999579e+00, -4.48875332e+00},
	{+1.79999578e+00, -4.28634357e+00},
	{+1.80999577e+00, -4.10057259e+00},
	{+1.81999576e+00, -3.92943788e+00},
	{+1.82999575e+00, -3.77124405e+00},
	{+1.83999574e+00, -3.62454867e+00},
	{+1.84999573e+00, -3.48811579e+00},
	{+1.85999572e+00, -3.36088037e+00},
	{+1.86999571e+00, -3.24191928e+00},
	{+1.87999570e+00, -3.13042974e+00},
	{+1.88999569e+00, -3.02570868e+00},
	{+1.89999568e+00, -2.92713881e+00},
	{+1.90999568e+00, -2.83417559e+00},
	{+1.91999567e+00, -2.74633670e+00},
	{+1.92999566e+00, -2.66319370e+00},
	{+1.93999565e+00, -2.58436465e+00},
	{+1.94999564e+00, -2.50950742e+00},
	{+1.95999563e+00, -2.43831539e+00},
	{+1.96999562e+00, -2.37051249e+00},
	{+1.97999561e+00, -2.30585003e+00},
	{+1.98999560e+00, -2.24410224e+00},
	{+1.99999559e+00, -2.18506527e+00},
	{+2.00999570e+00, -2.12855268e+00},
	{+2.01999569e+00, -2.07439613e+00},
	{+2.02999568e+00, -2.02244115e+00},
	{+2.03999567e+00, -1.97254586e+00},
	{+2.04999566e+00, -1.92458105e+00},
	{+2.05999565e+00, -1.87842774e+00},
	{+2.06999564e+00, -1.83397639e+00},
	{+2.07999563e+00, -1.79112613e+00},
	{+2.08999562e+00, -1.74978399e+00},
	{+2.09999561e+00, -1.70986378e+00},
	{+2.10999560e+00, -1.67128575e+00},
	{+2.11999559e+00, -1.63397622e+00},
	{+2.12999558e+00, -1.59786630e+00},
	{+2.13999557e+00, -1.56289220e+00},
	{+2.14999557e+00, -1.52899456e+00},
	{+2.15999556e+00, -1.49611795e+00},
	{+2.16999555e+00, -1.46421063e+00},
	{+2.17999554e+00, -1.43322420e+00},
	{+2.18999553e+00, -1.40311337e+00},
	{+2.19999552e+00, -1.37383604e+00},
	{+2.20999551e+00, -1.34535217e+00},
	{+2.21999550e+00, -1.31762457e+00},
	{+2.22999549e+00, -1.29061818e+00},
	{+2.23999548e+00, -1.26429999e+00},
	{+2.24999547e+00, -1.23863912e+00},
	{+2.25999546e+00, -1.21360612e+00},
	{+2.26999545e+00, -1.18917346e+00},
	{+2.27999544e+00, -1.16531503e+00},
	{+2.28999543e+00, -1.14200628e+00},
	{+2.29999542e+00, -1.11922395e+00},
	{+2.30999541e+00, -1.09694600e+00},
	{+2.31999540e+00, -1.07515144e+00},
	{+2.32999539e+00, -1.05382061e+00},
	{+2.33999538e+00, -1.03293455e+00},
	{+2.34999537e+00, -1.01247573e+00},
	{+2.35999537e+00, -9.92426991e-01},
	{+2.36999536e+00, -9.72772300e-01},
	{+2.37999535e+00, -9.53496337e-01},
	{+2.38999534e+00, -9.34584498e-01},
	{+2.39999533e+00, -9.16022897e-01},
	{+2.40999532e+00, -8.97798240e-01},
	{+2.41999531e+00, -8.79897952e-01},
	{+2.42999530e+00, -8.62309933e-01},
	{+2.43999529e+00, -8.45022678e-01},
	{+2.44999528e+00, -8.28025103e-01},
	{+2.45999527e+00, -8.11306775e-01},
	{+2.46999526e+00, -7.94857502e-01},
	{+2.47999525e+00, -7.78667688e-01},
	{+2.48999524e+00, -7.62728035e-01},
	{+2.49999523e+00, -7.47029722e-01},
	{+2.50999522e+00, -7.31564224e-01},
	{+2.51999521e+00, -7.16323376e-01},
	{+2.52999520e+00, -7.01299310e-01},
	{+2.53999519e+00, -6.86484516e-01},
	{+2.54999518e+00, -6.71871722e-01},
	{+2.55999517e+00, -6.57454014e-01},
	{+2.56999516e+00, -6.43224657e-01},
	{+2.57999516e+00, -6.29177213e-01},
	{+2.58999515e+00, -6.15305364e-01},
	{+2.59999514e+00, -6.01603210e-01},
	{+2.60999513e+00, -5.88064969e-01},
	{+2.61999512e+00, -5.74685037e-01},
	{+2.62999511e+00, -5.61457992e-01},
	{+2.63999510e+00, -5.48378646e-01},
	{+2.64999509e+00, -5.35441995e-01},
	{+2.65999508e+00, -5.22643089e-01},
	{+2.66999507e+00, -5.09977341e-01},
	{+2.67999506e+00, -4.97440100e-01},
	{+2.68999505e+00, -4.85026985e-01},
	{+2.69999504e+00, -4.72733706e-01},
	{+2.70999503e+00, -4.60556090e-01},
	{+2.71999502e+00, -4.48490173e-01},
	{+2.72999501e+00, -4.36531961e-01},
	{+2.73999500e+00, -4.24677730e-01},
	{+2.74999499e+00, -4.12923753e-01},
	{+2.75999498e+00, -4.01266456e-01},
	{+2.76999497e+00, -3.89702320e-01},
	{+2.77999496e+00, -3.78228009e-01},
	{+2.78999496e+00, -3.66840124e-01},
	{+2.79999495e+00, -3.55535537e-01},
	{+2.80999494e+00, -3.44310999e-01},
	{+2.81999493e+00, -3.33163530e-01},
	{+2.82999492e+00, -3.22090089e-01},
	{+2.83999491e+00, -3.11087757e-01},
	{+2.84999490e+00, -3.00153643e-01},
	{+2.85999489e+00, -2.89285004e-01},
	{+2.86999488e+00, -2.78479069e-01},
	{+2.87999487e+00, -2.67733127e-01},
	{+2.88999486e+00, -2.57044584e-01},
	{+2.89999485e+00, -2.46410862e-01},
	{+2.90999484e+00, -2.35829398e-01},
	{+2.91999483e+00, -2.25297749e-01},
	{+2.92999482e+00, -2.14813441e-01},
	{+2.93999481e+00, -2.04374075e-01},
	{+2.94999480e+00, -1.93977296e-01},
	{+2.95999479e+00, -1.83620781e-01},
	{+2.96999478e+00, -1.73302233e-01},
	{+2.97999477e+00, -1.63019374e-01},
	{+2.98999476e+00, -1.52769998e-01},
	{+2.99999475e+00, -1.42551899e-01},
	{+3.00999475e+00, -1.32362872e-01},
	{+3.01999474e+00, -1.22200802e-01},
	{+3.02999473e+00, -1.12063535e-01},
	{+3.03999472e+00, -1.01948954e-01},
	{+3.04999471e+00, -9.18549821e-02},
	{+3.05999470e+00, -8.17795396e-02},
	{+3.06999469e+00, -7.17205629e-02},
	{+3.07999468e+00, -6.16759993e-02},
	{+3.08999467e+00, -5.16438223e-02},
	{+3.09999466e+00, -4.16220054e-02},
	{+3.10999465e+00, -3.16085257e-02},
	{+3.11999464e+00, -2.16013715e-02},
	{+3.12999463e+00, -1.15985433e-02},
	{+3.13999462e+00, -1.59803370e-03},
}

//------------------------------------------------------------------------------
