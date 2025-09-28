package models

import (
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() *gorm.DB {
	dsn := os.Getenv("CONNECTION_STRING")
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database! " + err.Error())
	}

	database.AutoMigrate(
		&User{},
		&Role{},
		&Program{},
	)

	defaultRoles := []Role{
		{Name: "Boot Sequence", Description: "System initiated. Welcome to the grid.", Level: 1},
		{Name: "Ping Rookie", Description: "Sent your first ping into the void.", Level: 2},
		{Name: "Terminal Tinkerer", Description: "Exploring the command line like a curious cat.", Level: 3},
		{Name: "Packet Sniffer", Description: "Caught your first packet in the wild.", Level: 4},
		{Name: "Firewall Dodger", Description: "Slipped past basic defenses unnoticed.", Level: 5},
		{Name: "Syntax Seeker", Description: "Mastering the language of machines.", Level: 6},
		{Name: "Script Kid", Description: "Wrote your first script. It didn’t crash.", Level: 7},
		{Name: "Proxy Hopper", Description: "Jumped through proxies like stepping stones.", Level: 8},
		{Name: "Hash Hunter", Description: "Cracked your first hash. Sweet victory.", Level: 9},
		{Name: "Exploit Enthusiast", Description: "Found a vulnerability and smiled.", Level: 10},
		{Name: "Root Dreamer", Description: "You dream in sudo.", Level: 11},
		{Name: "Cyber Scout", Description: "Mapping the digital terrain.", Level: 12},
		{Name: "Zero-Day Whisperer", Description: "You know secrets others don’t.", Level: 13},
		{Name: "Darknet Tourist", Description: "Visited the shadows without getting lost.", Level: 14},
		{Name: "Keylogger Kid", Description: "Logged your first keystrokes.", Level: 15},
		{Name: "Backdoor Builder", Description: "Left a door open. Quietly.", Level: 16},
		{Name: "Shell Summoner", Description: "Summoned a shell from the abyss.", Level: 17},
		{Name: "Port Prober", Description: "Knocked on every port. Some answered.", Level: 18},
		{Name: "Encryption Breaker", Description: "Math bends to your will.", Level: 19},
		{Name: "Digital Phantom", Description: "You move without trace.", Level: 20},
		{Name: "Sniffing Specialist", Description: "You smell vulnerabilities from afar.", Level: 21},
		{Name: "Exploit Engineer", Description: "Crafted your own digital weapon.", Level: 22},
		{Name: "Payload Pilot", Description: "Delivered with precision.", Level: 23},
		{Name: "Trojan Trainer", Description: "Your code hides in plain sight.", Level: 24},
		{Name: "Botnet Recruiter", Description: "Built your first army.", Level: 25},
		{Name: "Kernel Patcher", Description: "Touched the heart of the system.", Level: 26},
		{Name: "Stealth Coder", Description: "Your code is invisible and elegant.", Level: 27},
		{Name: "Crypto Cracker", Description: "You speak fluent ciphertext.", Level: 28},
		{Name: "Digital Illusionist", Description: "Reality bends around your hacks.", Level: 29},
		{Name: "Malware Maestro", Description: "Conducted chaos with finesse.", Level: 30},
		{Name: "Sysadmin Slayer", Description: "Outsmarted the gatekeepers.", Level: 31},
		{Name: "Network Ninja", Description: "You slice through networks silently.", Level: 32},
		{Name: "Exploit Artist", Description: "Your code is both deadly and beautiful.", Level: 33},
		{Name: "Shadow Crawler", Description: "You live in the logs — unseen.", Level: 34},
		{Name: "Phishing Prodigy", Description: "Your bait is irresistible.", Level: 35},
		{Name: "Code Whisperer", Description: "Machines listen when you speak.", Level: 36},
		{Name: "Binary Bender", Description: "You reshape the ones and zeros.", Level: 37},
		{Name: "Digital Ghost", Description: "You were here. But no one knows.", Level: 38},
		{Name: "Root Raider", Description: "You own the system. Quietly.", Level: 39},
		{Name: "Cyber Alchemist", Description: "You turn exploits into gold.", Level: 40},
		{Name: "Protocol Pirate", Description: "You sail through forbidden streams.", Level: 41},
		{Name: "Exploit Evangelist", Description: "You teach others the way.", Level: 42},
		{Name: "Digital Strategist", Description: "Every move is calculated.", Level: 43},
		{Name: "Subroutine Sorcerer", Description: "Your functions cast spells.", Level: 44},
		{Name: "Stack Manipulator", Description: "Overflow is your playground.", Level: 45},
		{Name: "Firmware Forger", Description: "You rewrite the rules at the source.", Level: 46},
		{Name: "Quantum Teaser", Description: "You poke at the future.", Level: 47},
		{Name: "AI Hijacker", Description: "You bend intelligence to your will.", Level: 48},
		{Name: "Signal Scrambler", Description: "Noise is your ally.", Level: 49},
		{Name: "Cyber Architect", Description: "You design chaos with structure.", Level: 50},
		{Name: "Root Commander", Description: "You lead from the shadows.", Level: 51},
		{Name: "Digital Puppeteer", Description: "Systems dance to your strings.", Level: 52},
		{Name: "Exploit Oracle", Description: "You see vulnerabilities before they exist.", Level: 53},
		{Name: "Code Conjurer", Description: "You summon exploits from thin air.", Level: 54},
		{Name: "Darknet Diplomat", Description: "You negotiate in whispers.", Level: 55},
		{Name: "Payload Prophet", Description: "Your delivery is divine.", Level: 56},
		{Name: "Kernel Whisperer", Description: "The system speaks to you.", Level: 57},
		{Name: "Digital Warlord", Description: "You conquer without armies.", Level: 58},
		{Name: "Cyber Prophet", Description: "You predict breaches before they happen.", Level: 59},
		{Name: "Root Monarch", Description: "You rule the underground.", Level: 60},
		{Name: "Exploit Collector", Description: "Your arsenal is legendary.", Level: 61},
		{Name: "Malware Composer", Description: "Your symphony is silent destruction.", Level: 62},
		{Name: "Network Phantom", Description: "You pass through firewalls like smoke.", Level: 63},
		{Name: "Code Sculptor", Description: "Your exploits are works of art.", Level: 64},
		{Name: "Binary Prophet", Description: "You read the future in hex.", Level: 65},
		{Name: "Digital Sentinel", Description: "You guard secrets with silence.", Level: 66},
		{Name: "Cyber Nomad", Description: "You travel the grid without borders.", Level: 67},
		{Name: "Root Whisperer", Description: "Systems obey your quiet commands.", Level: 68},
		{Name: "Exploit Virtuoso", Description: "Your precision is unmatched.", Level: 69},
		{Name: "Digital Sovereign", Description: "You reign over the code.", Level: 70},
		{Name: "Kernel King", Description: "You sit on the throne of the OS.", Level: 71},
		{Name: "Cyber Messiah", Description: "You bring salvation through exploits.", Level: 72},
		{Name: "Digital Oracle", Description: "You know what others fear.", Level: 73},
		{Name: "Root Legend", Description: "Your name is spoken in hushed tones.", Level: 74},
		{Name: "Exploit Deity", Description: "You are the god of the grid.", Level: 75},
		{Name: "Code Immortal", Description: "Your scripts live forever.", Level: 76},
		{Name: "Binary Bard", Description: "You sing in machine language.", Level: 77},
		{Name: "Digital Titan", Description: "You shake the foundations of networks.", Level: 78},
		{Name: "Cyber Overlord", Description: "You command the digital realm.", Level: 79},
		{Name: "Root Emperor", Description: "Your rule is absolute.", Level: 80},
		{Name: "Exploit Saint", Description: "You perform miracles in code.", Level: 81},
		{Name: "Digital Myth", Description: "You exist in stories and legends.", Level: 82},
		{Name: "Kernel Overlord", Description: "You rewrite the laws of computing.", Level: 83},
		{Name: "Cyber Phoenix", Description: "You rise from every breach.", Level: 84},
		{Name: "Root Ascendant", Description: "You transcend the system.", Level: 85},
		{Name: "Exploit Eternal", Description: "Your code echoes through time.", Level: 86},
		{Name: "Digital Godfather", Description: "You built the underground.", Level: 87},
		{Name: "Cyber Sage", Description: "You teach the next generation.", Level: 88},
		{Name: "Root Oracle", Description: "You see all, know all.", Level: 89},
		{Name: "Exploit Immortal", Description: "Your legacy is unbreakable.", Level: 90},
		{Name: "Digital Legend", Description: "You are the story.", Level: 91},
		{Name: "Cyber Avatar", Description: "You embody the grid.", Level: 92},
		{Name: "Root Entity", Description: "You are one with the system.", Level: 93},
		{Name: "Exploit Infinity", Description: "You are beyond limits.", Level: 94},
		{Name: "Digital Ascendant", Description: "You are the final form.", Level: 95},
		{Name: "Cyber God", Description: "You created the grid.", Level: 96},
		{Name: "Root Origin", Description: "You are the beginning.", Level: 97},
		{Name: "Exploit Singularity", Description: "You are everything.", Level: 98},
		{Name: "System Overlord", Description: "You are the end and the beginning.", Level: 99},
		{Name: "Admin", Description: "Master of all levels.", Level: 100},
	}

	for i := range defaultRoles {
		database.Model(&Role{}).Create(&defaultRoles[i])
	}

	DB = database
	return database
}
