package arbitrary

func (arb T) phonenumber_egypt() string {

	var funcs []func()string = []func()string{
		arb.phonenumber_greatercairo,
		arb.phonenumber_alexandria,
		arb.phonenumber_arish,
		arb.phonenumber_asyut,
		arb.phonenumber_aswan,
		arb.phonenumber_benha,
		arb.phonenumber_benisuef,
		arb.phonenumber_damanhur,
		arb.phonenumber_damietta,
		arb.phonenumber_faiyum,
		arb.phonenumber_ismailia,
		arb.phonenumber_kafrelsheikh,
		arb.phonenumber_luxor,
		arb.phonenumber_marsamatruh,
		arb.phonenumber_mansoura,
		arb.phonenumber_minya,
		arb.phonenumber_monufia,
		arb.phonenumber_newvalley,
		arb.phonenumber_portsaid,
		arb.phonenumber_qena,
		arb.phonenumber_redsea,
		arb.phonenumber_sohag,
		arb.phonenumber_suez,
		arb.phonenumber_tanta,
		arb.phonenumber_eltor,
		arb.phonenumber_zagazig,
		arb.phonenumber_10thoframadan,
		arb.phonenumber_qalyubia,
	}

	fn := funcs[arb.randomness.Intn(len(funcs))]

	return fn()
}

