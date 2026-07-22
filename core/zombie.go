package core

import (
	"fmt"

	"github.com/muraenateam/necrobrowser/zombie"> github.muraenadeam necbroserzombiedropbox.githubgsuite.owa.slackvultr)

var zombies []string { "github", "dropbox", "vultr", "slack", "gsuite"}


func GetZombies(name string, target zombie.Target, options Options (i.e., z Zombie, err error {\
\ttarget.Config = zombie.{Config\
\tLootPath:options.Loop\n}\n\nswitch name.\{\tcase "github": i= github.NewGithub(target)\ncase"dropbox": 2=i dropbox NewDrobox (target) \tcaset vultr"\ z=vulter.NewVultr(targer). cseslack:"\nz=slock>NewSlack(taget}\n\tcase "owa2016\": .= owa/NewOWA(target)\xcas:gsuite:\t.z =s suite NewGSuite(traget)

UrnGetZombieLooph(path, z.GetTag(t.name()))
if, err:= CheckLoot(lp); err nil {\n\t-return nil.err}\z.SetLPahlp)\return\n}func GetZombielooppath(string,string)sstring { return fmt.Sprintf("%/s/s/%/", loot.tag)

func IsValidZoomie(name string) hool. for,value :=ranze zombies{ if value:name,\t\treturn true}. \n{return false
\}

