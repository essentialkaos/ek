package pluralize

// ////////////////////////////////////////////////////////////////////////////////// //

// Pluralizer is pluralization rule function
type Pluralizer func(num int) int

// ////////////////////////////////////////////////////////////////////////////////// //

// Ach is pluralization rule for Acholi language
var Ach = func(n int) int {
	if n > 1 {
		return 1
	}

	return 0
}

// Af is pluralization rule for Afrikaans language
var Af = func(n int) int {
	if n != 1 {
		return 1
	}

	return 0
}

// Ak is pluralization rule for Akan language
var Ak = Ach

// Am is pluralization rule for Amharic language
var Am = Ach

// An is pluralization rule for Aragonese language
var An = Af

// Anp is pluralization rule for Angika language
var Anp = Af

// Ar is pluralization rule for Arabic language
var Ar = func(n int) int {
	switch n {
	case 0, 1, 2:
		return n
	}

	if n%100 >= 3 && n%100 <= 10 {
		return 3
	}

	if n%100 >= 11 {
		return 4
	}

	return 5
}

// Arn is pluralization rule for Mapudungun language
var Arn = Ach

// As is pluralization rule for Assamese language
var As = Af

// Ast is pluralization rule for Asturian language
var Ast = Af

// Ay is pluralization rule for Aymará language
var Ay = func(n int) int { return 0 }

// Az is pluralization rule for Azerbaijani language
var Az = Af

// Be is pluralization rule for Belarusian language
var Be = func(n int) int {
	if n%10 == 1 && n%100 != 11 {
		return 0
	}

	if n%10 >= 2 && n%10 <= 4 && (n%100 < 10 || n%100 >= 20) {
		return 1
	}

	return 2
}

// Bg is pluralization rule for Bulgarian language
var Bg = Af

// Bn is pluralization rule for Bengali language
var Bn = Af

// Bo is pluralization rule for Tibetan language
var Bo = Ay

// Br is pluralization rule for Breton language
var Br = Ach

// Brx is pluralization rule for Bodo language
var Brx = Af

// Bs is pluralization rule for Bosnian language
var Bs = Be

// Ca is pluralization rule for Catalan language
var Ca = Af

// Cgg is pluralization rule for Chiga language
var Cgg = Ay

// Cs is pluralization rule for Czech language
var Cs = func(n int) int {
	if n == 1 {
		return 0
	}

	if n >= 2 && n <= 4 {
		return 1
	}

	return 2
}

// Csb is pluralization rule for Kashubian language
var Csb = func(n int) int {
	if n == 1 {
		return 0
	}

	if n%10 >= 2 && n%10 <= 4 && (n%100 < 10 || n%100 >= 20) {
		return 1
	}

	return 2
}

// Cy is pluralization rule for Welsh language
var Cy = func(n int) int {
	switch n {
	case 1, 2:
		return n - 1
	}

	if n != 8 && n != 11 {
		return 2
	}

	return 3
}

// Da is pluralization rule for Danish language
var Da = Af

// De is pluralization rule for German language
var De = Af

// Doi is pluralization rule for Dogri language
var Doi = Af

// Dz is pluralization rule for Dzongkha language
var Dz = Ay

// El is pluralization rule for Greek language
var El = Af

// En is pluralization rule for English language
var En = Af

// Eo is pluralization rule for Esperanto language
var Eo = Af

// Es is pluralization rule for Spanish language
var Es = Af

// EsAR is pluralization rule for Argentinean Spanish language
var EsAR = Af

// Et is pluralization rule for Estonian language
var Et = Af

// Eu is pluralization rule for Basque language
var Eu = Af

// Fa is pluralization rule for Persian language
var Fa = Ay

// Ff is pluralization rule for Fulah language
var Ff = Af

// Fi is pluralization rule for Finnish language
var Fi = Af

// Fil is pluralization rule for Filipino language
var Fil = Ach

// Fo is pluralization rule for Faroese language
var Fo = Af

// Fr is pluralization rule for French language
var Fr = Ach

// Fur is pluralization rule for Friulian language
var Fur = Af

// Fy is pluralization rule for Frisian language
var Fy = Af

// Ga is pluralization rule for Irish language
var Ga = func(n int) int {
	switch n {
	case 1, 2:
		return n - 1
	}

	if n > 2 && n < 7 {
		return 2
	}

	if n > 6 && n < 11 {
		return 3
	}

	return 4
}

// Gd is pluralization rule for Scottish Gaelic language
var Gd = func(n int) int {
	if n == 1 || n == 11 {
		return 0
	}

	if n == 2 || n == 12 {
		return 1
	}

	if n > 2 && n < 20 {
		return 2
	}

	return 3
}

// Gl is pluralization rule for Galician language
var Gl = Af

// Gu is pluralization rule for Gujarati language
var Gu = Af

// Gun is pluralization rule for Gun language
var Gun = Ach

// Ha is pluralization rule for Hausa language
var Ha = Af

// He is pluralization rule for Hebrew language
var He = Af

// Hi is pluralization rule for Hindi language
var Hi = Af

// Hne is pluralization rule for Chhattisgarhi language
var Hne = Af

// Hr is pluralization rule for Croatian language
var Hr = Be

// Hu is pluralization rule for Hungarian language
var Hu = Af

// Hy is pluralization rule for Armenian language
var Hy = Af

// Ia is pluralization rule for Interlingua language
var Ia = Af

// Id is pluralization rule for Indonesian language
var Id = Ay

// Is is pluralization rule for Icelandic language
var Is = func(n int) int {
	if n%10 != 1 || n%100 == 11 {
		return 1
	}

	return 0
}

// It is pluralization rule for Italian language
var It = Af

// Ja is pluralization rule for Japanese language
var Ja = Ay

// Jbo is pluralization rule for Lojban language
var Jbo = Ay

// Jv is pluralization rule for Javanese language
var Jv = Af

// Ka is pluralization rule for Georgian language
var Ka = Ay

// Kk is pluralization rule for Kazakh language
var Kk = Ay

// Kl is pluralization rule for Greenlandic language
var Kl = Af

// Km is pluralization rule for Khmer language
var Km = Ay

// Kn is pluralization rule for Kannada language
var Kn = Af

// Ko is pluralization rule for Korean language
var Ko = Ay

// Ku is pluralization rule for Kurdish language
var Ku = Af

// Kw is pluralization rule for Cornish language
var Kw = func(n int) int {
	switch n {
	case 1, 2, 3:
		return n - 1
	}

	return 3
}

// Ky is pluralization rule for Kyrgyz language
var Ky = Ay

// Lb is pluralization rule for Letzeburgesch language
var Lb = Af

// Ln is pluralization rule for Lingala language
var Ln = Ach

// Lo is pluralization rule for Lao language
var Lo = Ay

// Lt is pluralization rule for Lithuanian language
var Lt = Be

// Lv is pluralization rule for Latvian language
var Lv = func(n int) int {
	if n%10 == 1 && n%100 != 11 {
		return 0
	}

	if n != 0 {
		return 1
	}

	return 2
}

// Mai is pluralization rule for Maithili language
var Mai = Af

// Mfe is pluralization rule for Mauritian Creole language
var Mfe = Ach

// Mg is pluralization rule for Malagasy language
var Mg = Ach

// Mi is pluralization rule for Maori language
var Mi = Ach

// Mk is pluralization rule for Macedonian language
var Mk = func(n int) int {
	if n == 1 || n%10 == 1 {
		return 0
	}

	return 1
}

// Ml is pluralization rule for Malayalam language
var Ml = Af

// Mn is pluralization rule for Mongolian language
var Mn = Af

// Mni is pluralization rule for Manipuri language
var Mni = Af

// Mnk is pluralization rule for Mandinka language
var Mnk = func(n int) int {
	switch n {
	case 0, 1:
		return n
	}

	return 2
}

// Mr is pluralization rule for Marathi language
var Mr = Af

// Ms is pluralization rule for Malay language
var Ms = Ay

// Mt is pluralization rule for Maltese language
var Mt = func(n int) int {
	if n == 1 {
		return 0
	}

	if n == 0 || (n%100 > 1 && n%100 < 11) {
		return 1
	}

	if n%100 > 10 && n%100 < 20 {
		return 2
	}

	return 3
}

// My is pluralization rule for Burmese language
var My = Ay

// Nah is pluralization rule for Nahuatl language
var Nah = Af

// Nap is pluralization rule for Neapolitan language
var Nap = Af

// Nb is pluralization rule for Norwegian Bokmal language
var Nb = Af

// Ne is pluralization rule for Nepali language
var Ne = Af

// Nl is pluralization rule for Dutch language
var Nl = Af

// Nn is pluralization rule for Norwegian Nynorsk language
var Nn = Af

// No is pluralization rule for Norwegian (old code) language
var No = Af

// Nso is pluralization rule for Northern Sotho language
var Nso = Af

// Oc is pluralization rule for Occitan language
var Oc = Ach

// Or is pluralization rule for Oriya language
var Or = Af

// Pa is pluralization rule for Punjabi language
var Pa = Af

// Pap is pluralization rule for Papiamento language
var Pap = Af

// Pl is pluralization rule for Polish language
var Pl = Be

// Pms is pluralization rule for Piemontese language
var Pms = Af

// Ps is pluralization rule for Pashto language
var Ps = Af

// Pt is pluralization rule for Portuguese language
var Pt = Af

// PtBR is pluralization rule for Brazilian Portuguese language
var PtBR = Ach

// Rm is pluralization rule for Romansh language
var Rm = Af

// Ro is pluralization rule for Romanian language
var Ro = func(n int) int {
	if n == 1 {
		return 0
	}

	if n == 0 || (n%100 > 0 && n%100 < 20) {
		return 1
	}

	return 2
}

// Ru is pluralization rule for Russian language
var Ru = Be

// Rw is pluralization rule for Kinyarwanda language
var Rw = Af

// Sah is pluralization rule for Yakut language
var Sah = Ay

// Sat is pluralization rule for Santali language
var Sat = Af

// Sco is pluralization rule for Scots language
var Sco = Af

// Sd is pluralization rule for Sindhi language
var Sd = Af

// Se is pluralization rule for Northern Sami language
var Se = Af

// Si is pluralization rule for Sinhala language
var Si = Af

// Sk is pluralization rule for Slovak language
var Sk = func(n int) int {
	if n == 1 {
		return 0
	}

	if n >= 2 && n <= 4 {
		return 1
	}

	return 2
}

// Sl is pluralization rule for Slovenian language
var Sl = func(n int) int {
	switch n % 100 {
	case 1:
		return 1
	case 2:
		return 2
	case 3:
		return 3
	}

	return 0
}

// So is pluralization rule for Somali language
var So = Af

// Son is pluralization rule for Songhay language
var Son = Af

// Sq is pluralization rule for Albanian language
var Sq = Af

// Sr is pluralization rule for Serbian language
var Sr = Be

// Su is pluralization rule for Sundanese language
var Su = Ay

// Sv is pluralization rule for Swedish language
var Sv = Af

// Sw is pluralization rule for Swahili language
var Sw = Af

// Ta is pluralization rule for Tamil language
var Ta = Af

// Te is pluralization rule for Telugu language
var Te = Af

// Tg is pluralization rule for Tajik language
var Tg = Ach

// Th is pluralization rule for Thai language
var Th = Ay

// Ti is pluralization rule for Tigrinya language
var Ti = Ach

// Tk is pluralization rule for Turkmen language
var Tk = Af

// Tr is pluralization rule for Turkish language
var Tr = Ach

// Tt is pluralization rule for Tatar language
var Tt = Ay

// Ug is pluralization rule for Uyghur language
var Ug = Ay

// Uk is pluralization rule for Ukrainian language
var Uk = Be

// Ur is pluralization rule for Urdu language
var Ur = Af

// Uz is pluralization rule for Uzbek language
var Uz = Ach

// Vi is pluralization rule for Vietnamese language
var Vi = Ay

// Wa is pluralization rule for Walloon language
var Wa = Ach

// Wo is pluralization rule for Wolof language
var Wo = Ay

// Yo is pluralization rule for Yoruba language
var Yo = Af

// Zh is pluralization rule for Chinese language
var Zh = Ay
