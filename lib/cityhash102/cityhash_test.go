package cityhash102

import (
	"testing"
)

const (
	kSeed0 uint64 = 1234567
	kSeed1 uint64 = k0
)

type TestCase struct {
	key   string
	lower uint64
	upper uint64
}

var testdata = []TestCase{

	TestCase{"fa37JncCHr", 0x4aea51d5ee025184, 0x2d1433a519898c72},
	TestCase{"yDsbzayy4c", 0x8aa1ace28e1aed35, 0x795edb57643ec05d},
	TestCase{"BWDxS22Jjz", 0xc86504d7706144d2, 0x6b0b14cbb8e6a889},
	TestCase{"hMaiRrV41m", 0x17789d956fb03bb1, 0x4eaac081d89aa328},
	TestCase{"tzxlYvKWrO", 0x70e083476eb90747, 0x7d4662c570540e4e},
	TestCase{"72tK0LK0e1", 0xc7f3844273674ba2, 0xeed8e3defff28a2a},
	TestCase{"zLOZ2nOXpP", 0x8e3e01c38d8a3ffb, 0xb8c98ad62081d551},
	TestCase{"IhMFSv8kP0", 0xd0d32a0c4ee1aa6e, 0x1796ec90d57c9b05},
	TestCase{"7U20o0J90x", 0x86be1f3d9f4a90af, 0x3a2f121c710f2cb0},
	TestCase{"A0GWXIIwo7", 0xad526ea1e1a79ea, 0x93b529c61611a5ea},
	TestCase{"J4ogHFZQxw", 0x57ca60e310b63fcc, 0x3eed6c0e19ff8a53},
	TestCase{"Q2RQ0DRJKR", 0xe106311a63027535, 0xada99e8cc914290c},
	TestCase{"ETPVzxlFrX", 0xaaf12e879ba92174, 0x55e1beb6f791ca5a},
	TestCase{"L8b7mtKLHI", 0x152f1f95a626be95, 0x3a4aab6dcf2b652c},
	TestCase{"GhIh5JuWcF", 0x742b013f921beba6, 0x48e6114d7c53995a},
	TestCase{"wrgJKdE3t5", 0x6af6210249c851bd, 0x82a47d925824cf1e},
	TestCase{"bECALy3eKI", 0xe385b545102e28f, 0xbb628b3c3c8a24a0},
	TestCase{"wYxEF3V7Z8", 0x20d7c715fb5c4ab, 0x5dacb3afa7831557},
	TestCase{"KTx0nFe1IX", 0xcc7cefdc2acd9e91, 0x8dfe8f853922e929},
	TestCase{"5tjH22F5gX", 0x298d9bbbcb9b8e76, 0x26890d91e245485f},
	TestCase{"Oa5LnIMIQu", 0x51e43a89c83ab206, 0xf03667e163c5c619},
	TestCase{"OiNJj8YL8r", 0xe428ebb3d807a831, 0xb2f39e8cb204a38a},
	TestCase{"qDiZSkZfoE", 0xc43ba8f23fbff6e6, 0xbab1358363a3e88c},
	TestCase{"DAmGTXXqqv", 0x783e1b03d89f0555, 0xa55d1826dca7d29},
	TestCase{"kCd5WKE2fM", 0x6a3c67d38392209e, 0x1e31d94c82a004a9},
	TestCase{"tVXa2zKae6", 0x8a4004fa751c96af, 0xfb5573a5ad97658f},
	TestCase{"opGY4i6bYu", 0x5d944005b98aa22b, 0xf3de24c917343988},
	TestCase{"UG67LaSXd5", 0xbeffb339af07d9d, 0xe246dc7823a1ecf0},
	TestCase{"tUbO4bNPB0", 0x83bcecf731764f6e, 0x2e49adeccce3a6bf},
	TestCase{"TxnkWrSaQy", 0xb4ac06c7f5da386e, 0xdfd55c8f2aa6141f},
	TestCase{"fa37JncCHryDsbzayy4cBWDxS22JjzhMaiRrV41mtzxlYvKWrO72tK0LK0e1zLOZ2nOXpPIhMFSv8kP07U20o0J90xA0GWXIIwo7J4ogHFZQxwQ2RQ0DRJKRETPVzxlFrXL8b7mtKLHIGhIh5JuWcFwrgJKdE3t5bECALy3eKIwYxEF3V7Z8KTx0nFe1IX5tjH22F5gXOa5LnIMIQuOiNJj8YL8rqDiZSkZfoEDAmGTXXqqvkCd5WKE2fMtVXa2zKa", 0xa746b7532685199, 0x8ef5fb78800ffdbd},
	TestCase{"e6opGY4i6bYuUG67LaSXd5tUbO4bNPB0TxnkWrSaQyUuEa0X9Q5mVwG4JLgeipeBlQtFFJpgHJYTrWz0w2kQw1UFK8u2yWBjw3yCMlqc4M3tt2un4cDzdiEvq8vmf7TZAPjUAZ6Cu86nAyYDamCCSQ7GX33A8WhGwRk40pHuxNf5JEItyS3QrBgOChWKCDa6eIAd7RV4mBA5NQxJt0jk9N6L5cdFnDLSWV3bvYghhol4EgN5e4poSt7VVlkJw5jSYm", 0x2eef91b0bc55fe00, 0x1e7bedd30d4ded09},
	TestCase{"4TKi92Ws4iYQoCSbysV6Nyp5Fl8wCfiE81uF1O736dRsouSmmxq8tfB7PK3Zzmn5lhLm5Qn92F2q9UatPR1G4DNRVR0SBlXwQqgTFRdHgd5n5ffS4gi9r6YKVZmgIIaj8ECLfncKQh5TLkvPPcYEg5ZBeJpubNdiZq3CbeW2JcTeKP4j1ayffXqHqdCQ0n8Xb9jDnEF7oij85ls4MqjzLXF9APZ8CffopP1adEfRuPX0AP2UDmSWHhgS6DaIrE4eb5", 0xf16ca0b95b1deda4, 0x78fdcd68cfb1e4b6},
	TestCase{"EEJudCHACPYCulwMIE1wg57ENyQSc1VpFnjqz019PZLHIIbYWaSAfaM3WnT7oyw2jdsibrryODEhTpFzQi73GT6kGXr5Ul7DOxwxplwDyAuRx8OLoVP2zTmDzeITNNekLYh8KbLIjEihK408aNAXrwkoY1HwMtgfSLnmx72gLiLfnKlLhtsWpaKMZZGwTubvFNhAUhppQASDSBYA4OetwzDWYTQzNzubMZlqHadfj3sBEOJIkyAevNATpYRAYLlutV", 0x2941ebc55891e387, 0x94cedd6ef5685a8c},
	TestCase{"j85MnoOfyc1HvlF3N8QYaD41OcK7VDcELgY8SwlQXmiQVvTt4rPe5RdR4xYXB9lUpHdHCMgj7O7aHaRJRovWGYvKUUrfba7Qpif15LiChpkxNCGp0AJGgFYAhPnIxvgndJmgfTqKGbHenWRlgk2KxaVeyGuv9YinsTRVwIpCt7qedHPH0Pbx04awLSrS1YFr1fMvx97oGwQrBp89Di5Bmf757yY6UlvTQHOLRU9fQZXZNdhYLmj6RqBWmhbHRWkrm9", 0x23016e0814ffe84, 0x2d0e4abc3f3dd376},
	TestCase{"BBbIqzqLYDzFjK1SQQIav2HWJi22Ym9jxkzojp7F06TjRUBptRPoUfKlLKnr7uY2eYqLNwbO247RWHHNieBAHTwdohUtc3vEbkYyg9KiBS8fjP3P1EYJiUwU9ONjRGw00UxgbHNmjVRQsUotjMAPo4txTEfsUbrT3o9e5UQnxpBnIzfzLpO9uF5LTiDvH4OKqWywyMhw9sjRsOQBCmL61ORS6cONfmhVGdPFx6B4xsWpFu0RhJVihu9nWX89HndWQ2", 0xecbeb78c55867162, 0x8a3007296256a92b},
	TestCase{"lL7uQ4mutzmrQT9tAqnJcIoiR3W4Zw5KGCCeExW5wIwLm5Euu2BUKzCj0ioadrsr15VF21KwHEH1KWvCY6es3qb2XNa8CSzDVUSVTmSi1jkJFfTlk7blvBlSYLajmXwHzNlS7DB6utP8WqtGvV0rglHC5qtrNq6NBrnI2wPxpm3MbuaWPYN3KfEPT5EqtKB4CzNEtk9jWC377dcUCzWUcir5n5uhP5jZ26mwovdJ6gOBH0dR09tapkebnWGBZzkc4W", 0x919a0d94f6d70db2, 0xaefc152f63848bf3},
	TestCase{"WsQ9BWnah2aKUYSN7F6lqtF62o6nO8Hu0h1CBmkspKBJrbeyqkhfcwjekpP9vf1wM36YpqOc8Xvz8YvxIccsWLVH9sO4XNyuO4QCuoo3Ki0SGwNWYxP5JsKQkgSG3R8QVYcOKQRc8R5ONQvuOIxf8H5qvV6wwEMQkyo3PFfVeiu09nuV3rAB6F1zk6vfII51Gt2d6f9kO1iVocyrR80VL1U35QiOglPweRai6hTSi9vWmtLB1LfKK9OQX6mDsB8Uag", 0x571dd95c64d9abef, 0x27708ecb26bd6cc4},
	TestCase{"DgNe75n0ZXuujtDMEXVcv1ghCYqK1Q0E5NqAQbAz94tqw4CAafkVeQBrxz9yQAATV0dubltiqlYkpiuPMet042r111xR97sc5TWfFON39unwahMwJDxLFmNGlIguQYVV01AFNWIWO5Tyknv3yqODcjRO27GQelvcl4r8Y7dwC4supnxndJ0E1Pa3VqT7aMjLOaRuh2qr6jlvUgg6xeIy4szZiQgInPb9x14c3sT8ZC14si8pKQmOIlvyCZFxwq4tp8", 0x935d576c38ba00a7, 0x510fcf130aa2115b},
	TestCase{"VrzwxY8yayeioz6YLOJHLToaOlOTcEZ65Y2259yf7dMtaQRtnjA8Cxibi62IKbON8PNDYMqdxAXVbwOOfXWsUERAITSd2ryAEMLkgCNdKs6vpUHU1nKTzjdIA6tDvpN7BjrptCQD5W9szOK0CdS9M5PX9Gk2657Fmy5f9TrE0y4zMOxYzPgMT3raKZbOeidPhg4oAv28t47GQ2mQQQmtUbRmA2AmllBQQEEY7EgzGlFglz6BOq2sRUeZWmMFXVdxjp", 0x4448950711a661bd, 0x4adfee5763ab86e},
	TestCase{"Tr19oHs1ye04nMvqEKKquoaE28jf3RVXGUe4mV5k73muQhkc02Rur18t7pW9EzeUUGYGlbyqfkl3STfSW4KL6QEBDiKSiywAFTQy2Oph8YiY1MyXQGqUg4gsl0KRxGbAhz8kNxPTT80VUw0uBrPpt3he315yGe8xdHhyC5SgDQ9hK9ZV0wLsz2U14azIE8FqNvoX0EBDeKszSPSQMnILonKsLKAXQNNlICGGQSR3kJ0ChSd3GvM2ghs1z0ZNNu6e6M", 0xf9d5cc824bef681b, 0x7ee3565e87c44a57},
	TestCase{"uVoKYZbYjI0MLEGiGwN8wK8ThUNl8U70zuKXRw3AE3XYGlEWhbdbvj4aDSLLuQKrKeOlYPunSRLgAZBp8oR1XVZlvs4pIOgd12MZRGMsffYpChdIV4H1ZqkTioIzBzaC1wjRB4JojpcuXFC0JS1qIjJzWcyhZYrYSaxbeEQM42GZHQXYqXO9Gg6mG2RpaGL0rJcTX2pZ236JSerIBDPRrVC7Xbu5sE6jXgC2g0ci1i1TMqkV19vtd7y8irEa5IHcwS", 0xe7b664e616f50389, 0x41b26cbf08df4a96},
	TestCase{"ecQEKRxLtH9dn8kgzNnyWTpi2u0HUwj8Y7LqZIBQXI3IQlyQ9jMdB9LD3JUXDCdlJybqEkGk0H0O3zmAg8nrH64IPXnchPLyLupYc3IaJIyKFlSwrFn7JqPiNBI4ab0vVpR5rjf80bQDMs9C7vJQlg66rM8Px8JQwkVnU8tSjJf4BnGGiZeTFiZ54hU1pnPjYsU00NShe5lpsz3aYh1mPYrRDJQ37pkdfFdfc3LH844z48ZcnZMD7DcLX3Mcq6FWLq", 0x18df8365c68cafb9, 0x553bc712aa1cf699},
	TestCase{"9wtUDzZFydNWDB5aMClyVGzrqqw5MHvVDm0OmXek8zGLAJtUUfRztQpkGlnc2h6ET4aFbEyjCC2KWwp0ZGxQemAtXwTYdZk6dKKCXIvjSw3yqqwN4snjewcAs3gVdObGgtRDBKueGwa6mVTpNFYzB871BlWm951ozSz8muk1oJ7aoZP9mv6x1CwBvSv3VwrSMqb9iJAWcF5OmSXYMcTNmPWhpQkLMbnjSMq8gyfIBkhyCCUYmyuYLOGBoyU9aFq0cg", 0xc082a5550a38625, 0x587b4906db788f7d},
	TestCase{"8GclYnVDjiOEEAC8hVVvgJtASRQIR0wYGYJmKp03VMHhURoBxJ6bayk3P8Jp8ENMmg84T68wSNcxnQ6khBJH93IYBaNHoidYNkdqqllH6NCtlGbSRvjaw078aSNO9zwWjZKXI3mPQyGBDsbenLDjLIpviDHrCDNtngQ3jBS87jHKZtzMCA3VSrOA2f2CsN5d3Thne7vkoC4N33jFDlkda8ocnooeBrHDIw0v4tfs3jG7lxMwg5aGDMszAedLTsYkpW", 0xd71e4b47a89d52d2, 0x43770ea94fbf0820},
	TestCase{"ftNIjRzzWkvqfbvFs7Zi4jLh2mXZWM59DRPvqPuM7nAmP4zF9YvDFGsI1PrVkufvJ2q9RjTZWcJtgJ8pp425IuLJHCp14SvOUlVwSxTzZmqF3z2t34wLwHfDRSDTu8pOrLIIGlFpW54Z44Q58MR2b4D2UQUOWJkMc3esmrfIwhq1lG6rbVtCX7D1VfN0y8KZ9wPvn5Cjm2iYGnPrGH2oODnrsBrqJ9NS5nLsqVacVI8m3YbKnb69ot0g4qVNzsn5d7", 0x59593d66cd8705d3, 0xe039d3d955a1c272},
	TestCase{"xTcX66oCsriR9V3Fep9fWDT1ZSrKXVPTy2y28mFycxNlTO17E8kiLDjsgaBB3ac2ab2iLHevD0fgOglamUI5h1yLa7Vdh7dIggy0vbt8bYmxCWYy0q4hq01O7Wzobd4GH3GCc9JDg58sbgrbVvHLvIh0mgoNHrdYsrjUz2ff5nVhUMGxFVG8lx8XdvtvmUReKA8JCnwGYRx0mBx1fDAQ8GwkBndxF4ZZDhqNSmdzCYzykxxP85nFMjzVUcRke0Jrg8", 0x6b14829a48b6e85a, 0x3ff210965162bfe},
	TestCase{"F6ss54R5090yW63HLPzIvTsKDXKVM0dZ5TSAYrDWyDUSIYjdviuoCm8NHRsePTDUvdfRSqxQ2SsKyavrIpgScmqtDgVd8j73KkUnZRDbr4vpeqfueLNE7D5Jtbw1I12ckWNJwZsndnbFbG9DZWRehXvb8raOrbyZ7Mq1vinWTOl4csI9Phl6EhfKWEiNpfuu1kwwQjSt5Byf3EmQvXV8E8QlM769m12njyhAf71iIzLJB7k5eFBrLbcgigpSfrDPnu", 0xf2a3aa84c08027e2, 0x6fe66b1e59dc1c72},
	TestCase{"XQ1Y8HWTbfYJiDWt4rTeX9KLc0ApNwksMjOTzuuaYTrGeNAgCdKkmc3MaDCv7ulUbAva2pAaI1ruMzaZasHKTKh3Vrxdli5Mq0vtp5R65G0PFbynTD8wVny1ctcLZhhQhaHUgiajzb6CC30dG6YmrUlTLNou4TIj4ZDiIoPHNVTXYRAnXgZM8KpTfdOj7gQ9GdrWRGlolCJIeU598eWHwLkbx6I2niB1J2VjGGVzSpH4JMEPyigs3RRyVj0IPBHjBm", 0x1c51f60bdbbe9dbb, 0x39acacd69aab5b4d},
	TestCase{"SR2vRTkiX13lRzT7sVWJT21SIQdY7nKXELTdoDKLFN5gMWlCzGWRGXtYvU61HPYUkz7WDRpQmu47PpJO3nnJIgsEBwDQLlu3j0ZuPPKAHOHgBY3DMqWeXOqgJ36enzgWzFQOeiYw4naDMdQgSxIzJ7dcAhFvgvRf8p2lXYfbMFpgrFKJCcGUjt6sYJnFCDsL2s6aRl9lywRnCj6MLKq3BvvkCgzPtpiuioS9XbsTYJHi0L2LgtMpmFZyuWLnM3f2p7", 0x6c1eaa80d8db2afa, 0x100a372d6ea3c1ac},
	TestCase{"BKh3oDK3vKMxd0oyqYDNX5ustEuWGjeP4JQqWirP0DKbE9Z4flPCqJ3hVxEjeq8i7WWdEN0EbKqnRPp78EHzWIe1FsjsgrYnM5OaQOpzhdm60ZD9oT6KlkJzb2rFsN3ESRotobqVCdZDClKyEQIx8auhcjxU7ygZNVQB4HeHsES2zm1BAH8Gp3ySmtutrbRC4pM640LwCnx9ZyIjDPy3QuTAnL4cuTowGA3L9OHJCCRj8jRM6PNWJqh69jg4CS0Sd3", 0xeb7415a20c66fd9e, 0x8c3d8bceea00e4f1},
	TestCase{"nkQ31aDSJMAkgG73nQrSU0BA4Nb3nC4ZwUav4nOM7W4nkCqVcfx5e8DiTojF0oouIMqL7ChFil2SxrwYWtdA2orVdakdMWWfqMyyYfBEODeK2aqXURfUFUzs5iURC043K2zqfA42LiMMHCtjbYDq3Ah6q9V37Y4RY1HE9JGSzaoGkhzLEA9FLqJBzmC4KGVtHk5O3Jr1tfHdLFwXP5miu3urp6wALP1a84z9NqAEVQpqflN2o7liAfayjU84t7f1Bc", 0xe87302b12e25c7ea, 0x478b39e3610b0d80},
	TestCase{"BWSLkxjZmOK7P6DApLnPJUtPZkWCmfovBEEt9ySvKk1jrEteXe3r9uGgclqMQcHaqURxQutkcsUR4L5by6S5yjlbSZxqBDOzhqw7inpKdJmhdrJZwjeuQPVswTj5e72LxwQfkFyNXi38XKhT3LOUitKeK1hy9jH4fhiNvelQMmYt6DN9WlbEcwqxvXv4ED7tsnHlR0BlmhequzzQiZcLTQIOvBR9MW1DJIyiG8T2p5sk3p8lMj6p9MC5WdCs9D3QT2", 0xabd5a76e78640f9, 0xe02191d1edbf17b8},
	TestCase{"6jAXjxaZhdPoPlXTZeqlhMMsCT5GtZGzGOU0L4Z1iwn5gIYFwOyciKUtnX7e6NbLl5L49u5ppquV8Tk5pifW09PngWPlr06a3PeCHi06YubhNJkBzPhzW6kAa9tR7x1AMfLcLLiuFIZbbHmbgTaAXJK7SEYYBZiWC18YMqQa8zBiExJsOr2vAK2bYY9i7pCHrKpD8DnFAwxOrEGF5I8Gb8p7gypll13cJqoS2ZhCWcbLqpbv6h9hoyoUucEgbFGv34", 0x66fe27da6e49bf24, 0x39b17d86062d6733},
	TestCase{"N5e2I8erRUg0PkiXRWVEyOqC4RPKMTNhYzjo5LfWpKXD4FkUjDihZWsdxHwIkJxHGf5L0jro1O13bjXKuD0Ujs5f71xqIu7YZAtYtjMu7Nwh4T1wfzQOrU3zTypmtuIQ4AyytKqyfkfhEfcte0FWUHTyFGi8A0XCAT83mx2Phh5vKhmzh0TCHwAWAsdIq8U1bd2Na4lHkqB4Xv1EvUOAQWhYMIqDQLEzyELYJ6n1wy3Rs5dnZ2wyYdWvtN6HiIGeWb", 0xe64592f5f4e7704c, 0xd3830d79326a0779},
	TestCase{"Anfxobuq0mveXSgTQC4v5xGBEwTTa04inifaHXQIJJwplaIAmN5qIL1VFTypR0VCiAm0gAIxTCmDm5NWQQMijMByp9lGAGSqODq4N60pIm22pNZDltuUF5Q2FBGNPjDnu4qFAq2Ra2TQP0bAuVc7b2AoBQBa7OOzQCFZ2F0cIT0fUcqO5QTgTbSc0dD70Z5OlIxlXxNpRMVvwJH1jigAK8lKkyPiXU6Gm41I1O7QiaJerYeZHKjbQUv8QIqvnxBXzC", 0xad3c40aa10603113, 0xb9e85851669951a7},
	TestCase{"p0ZvPFTiuLHWuYqb7E40NUIBQ38Zb5kO6JHTYlBQ6Qw0ykb3xf3K8MVWNc5yhpLl6cEdLO3RoyPlgzodcsxkCRGZ1KWg7rSERern2sCpocaVbO8DE4wQTAyUSUAaJamiFbVFRg4GGclrzr2DvweO5cqX417NZr5mSZ2uD4ATftIekIreFT0K3rp7qxSPmWCF3C9GEHht8zXsFNUUqVmsMczCXQZJulWxvdBAsq30pZr4wLXkoHcAraKM0uesdAnZmy", 0x2acf14b2bf64cf07, 0x9ab3b1027f4a145d},
	TestCase{"hemkfbHWdBp8weQYmF67c4zGwaOj7AfopRXU0oyc0lkuP8QAOXHybEEWocFwltkaKF2J4yv2jdw7mMF8tW6TiIzUsCOd57BOMBhQ9aSrDMwxjB3cf75LP2pGFEtKJ3gfEN4MvUD7r74aI6AvDEGbG5rTHioalUGyqIKjmVqdctBuzMpAa5lo8cIQL4y4WC2KUK4HpssRl4JiQ6ty9ckIF2gY4eabrbtJvxalqTCZXWGva9tklb2yciUgK3F9e9TX41", 0xe750c01b24e1ebaa, 0xe4d7220d36ee2151},
	TestCase{"IsUTSzxiuXplHYLJVxzxdI0rPcysA2sSvMtLKq1COowdKHvpEslpAjeaMcQWfIwYdorveq70c1ewGXkSPVIYCu6YXW2AnxiOlXJNNOOzQ0tgXd9u6PSJJYpo5syqngEYBVvYtHVHFPyk05f7S7OleDYh5WXsBjOMEIt5XONnlJVkOArqGDZuQ7bUe6Knpj91z06WNTJ8mpsAxhyBuW5JbenDk70XoAYnAfKX8betOU1KA0V2UbL5D6JwEJT0R0oaf8", 0x6bac4aa1415bf757, 0x29fe5e81561c3623},
	TestCase{"7lhje6CeQMevO6UiBgoUc0o31F3npitvSac4eoi2AKwXQQFc414g2sk16nmtUfmuFPysDfuMzqrOF6yJ5205vi4zTqsxVdrj2obDTTXSKOoXSmqWpobiUghyWXv0AlhAZIM2jrS1GHYg1PCqBmYgSDewlXutIb2prMrbDIcRXB8YYIPh4vNU9zPsXJjnvldk5ULJkwkH5qqd9DKB8ggHf39CMt0Fcbyi5JzoDj5IYtvf6EoFsSUXVciqVg55H1lKLk", 0x808b0f6e328ac24f, 0x9fa199beefcaa7cb},
}

func check(expected, actual uint64, t *testing.T) {
	if expected != actual {
		t.Errorf("ERROR: expected 0x%x but got 0x%x\n", expected, actual)
	}
}

func test(str string, lower uint64, upper uint64, t *testing.T) {
	var u Uint128 = CityHash128([]byte(str), uint32(len(str)))

	check(lower, u.Lower64(), t)
	check(upper, u.Higher64(), t)
}

func TestHash(t *testing.T) {
	var i int
	for i = 0; i < len(testdata); i++ {
		t.Logf("INFO: offset = %d, length = %d", i, len(testdata))
		test(testdata[i].key, testdata[i].lower, testdata[i].upper, t)
	}
	return
}
