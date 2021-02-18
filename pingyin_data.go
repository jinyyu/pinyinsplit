package pinyinsplit

//https://baike.baidu.com/item/%E6%8B%BC%E9%9F%B3%E5%AD%97%E6%AF%8D%E8%A1%A8/5428784?fr=aladdin

var pingyinData = []string{
	"ai",
	"ei",
	"ui",
	"ao",
	"ou",
	"iu",
	"ie",
	"üe",
	"er",
	"an",
	"en",
	"in",
	"un",
	"ün",
	"ang",
	"eng",
	"ing",
	"ong",
	"zhi",
	"chi",
	"shi",
	"ri",
	"zi",
	"ci",
	"si",
	"yi",
	"wu",
	"yu",
	"ye",
	"yue",
	"yuan",
	"yin",
	"yun",
	"ying",
	"b",
	"ba",
	"bo",
	"bai",
	"bei",
	"bao",
	"ban",
	"ben",
	"bang",
	"beng",
	"bi",
	"bie",
	"biao",
	"bian",
	"bin",
	"bing",
	"bu",
	"p",
	"pa",
	"po",
	"pai",
	"pao",
	"pou",
	"pan",
	"pen",
	"pei",
	"pang",
	"peng",
	"pi",
	"pie",
	"piao",
	"pian",
	"pin",
	"ping",
	"pu",
	"m",
	"ma",
	"mo",
	"me",
	"mai",
	"mao",
	"mou",
	"man",
	"men",
	"mei",
	"mang",
	"meng",
	"mi",
	"mie",
	"miao",
	"miu",
	"mian",
	"min",
	"ming",
	"mu",
	"f",
	"fa",
	"fo",
	"fei",
	"fou",
	"fan",
	"fen",
	"fang",
	"feng",
	"fu",
	"d",
	"da",
	"de",
	"dai",
	"dei",
	"dao",
	"dou",
	"dan",
	"dang",
	"den",
	"deng",
	"di",
	"die",
	"diao",
	"diu",
	"dian",
	"ding",
	"dong",
	"du",
	"duan",
	"dun",
	"dui",
	"duo",
	"t",
	"ta",
	"te",
	"tai",
	"tao",
	"tou",
	"tan",
	"tang",
	"teng",
	"ti",
	"tie",
	"tiao",
	"tian",
	"ting",
	"tong",
	"tu",
	"tuan",
	"tun",
	"tui",
	"tuo",
	"n",
	"na",
	"nai",
	"nei",
	"nao",
	"ne",
	"nen",
	"nan",
	"nang",
	"neng",
	"ni",
	"nie",
	"niao",
	"niu",
	"nian",
	"nin",
	"niang",
	"ning",
	"nong",
	"nou",
	"nu",
	"nuan",
	"nun",
	"nuo",
	"nü",
	"nüe",
	"l",
	"la",
	"le",
	"lo",
	"lai",
	"lei",
	"lao",
	"lou",
	"lan",
	"lang",
	"leng",
	"li",
	"lia",
	"lie",
	"liao",
	"liu",
	"lian",
	"lin",
	"liang",
	"ling",
	"long",
	"lu",
	"luo",
	"lou",
	"luan",
	"lun",
	"lü",
	"lüe",
	"g",
	"ga",
	"ge",
	"gai",
	"gei",
	"gao",
	"gou",
	"gan",
	"gen",
	"gang",
	"geng",
	"gong",
	"gu",
	"gua",
	"guai",
	"guan",
	"guang",
	"gui",
	"guo",
	"k",
	"ka",
	"ke",
	"kai",
	"kao",
	"kou",
	"kan",
	"ken",
	"kang",
	"keng",
	"kong",
	"ku",
	"kua",
	"kuai",
	"kuan",
	"kuang",
	"kui",
	"kun",
	"kuo",
	"h",
	"ha",
	"he",
	"hai",
	"han",
	"hei",
	"hao",
	"hou",
	"hen",
	"hang",
	"heng",
	"hong",
	"hu",
	"hua",
	"huai",
	"huan",
	"hui",
	"huo",
	"hun",
	"huang",
	"j",
	"ji",
	"jia",
	"jie",
	"jiao",
	"jiu",
	"jian",
	"jin",
	"jiang",
	"jing",
	"jiong",
	"ju",
	"juan",
	"jun",
	"jue",
	"q",
	"qi",
	"qia",
	"qie",
	"qiao",
	"qiu",
	"qian",
	"qin",
	"qiang",
	"qing",
	"qiong",
	"qu",
	"quan",
	"qun",
	"que",
	"x",
	"xi",
	"xia",
	"xie",
	"xiao",
	"xiu",
	"xian",
	"xin",
	"xiang",
	"xing",
	"xiong",
	"xu",
	"xuan",
	"xun",
	"xue",
	"zh",
	"zha",
	"zhe",
	"zhi",
	"zhai",
	"zhao",
	"zhou",
	"zhan",
	"zhen",
	"zhang",
	"zheng",
	"zhong",
	"zhu",
	"zhua",
	"zhuai",
	"zhuan",
	"zhuang",
	"zhun",
	"zhui",
	"zhuo",
	"ch",
	"cha",
	"che",
	"chi",
	"chai",
	"chao",
	"chou",
	"chan",
	"chen",
	"chang",
	"cheng",
	"chong",
	"chu",
	"chua",
	"chuai",
	"chuan",
	"chuang",
	"chun",
	"chui",
	"chuo",
	"sh",
	"sha",
	"she",
	"shi",
	"shai",
	"shao",
	"shou",
	"shan",
	"shen",
	"shang",
	"sheng",
	"shu",
	"shua",
	"shuai",
	"shuan",
	"shuang",
	"shun",
	"shui",
	"shuo",
	"r",
	"re",
	"ri",
	"rao",
	"rou",
	"ran",
	"ren",
	"rang",
	"reng",
	"rong",
	"ru",
	"ruan",
	"run",
	"ruo",
	"z",
	"za",
	"ze",
	"zi",
	"zai",
	"zao",
	"zan",
	"zou",
	"zang",
	"zei",
	"zen",
	"zeng",
	"zong",
	"zu",
	"zuan",
	"zun",
	"zui",
	"zuo",
	"c",
	"ca",
	"ce",
	"ci",
	"cai",
	"cao",
	"cou",
	"can",
	"cen",
	"cang",
	"ceng",
	"cong",
	"cu",
	"cuan",
	"cun",
	"cui",
	"cuo",
	"s",
	"sa",
	"se",
	"si",
	"sai",
	"sao",
	"sou",
	"san",
	"sen",
	"sang",
	"seng",
	"song",
	"su",
	"suan",
	"sun",
	"sui",
	"suo",
	"y",
	"ya",
	"yao",
	"you",
	"yan",
	"yang",
	"yu",
	"ye",
	"yue",
	"yuan",
	"yi",
	"yin",
	"yun",
	"ying",
	"yo",
	"yong",
	"w",
	"wa",
	"wo",
	"wai",
	"wei",
	"wan",
	"wen",
	"wang",
	"weng",
	"wu",
}
