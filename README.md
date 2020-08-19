# godartsMySQL
This is a GO program using templates to build, update, show and edit the darts.player_stats MySQL table.
SQL to create and insert test data:
BEGINNING of SQL 
{\rtf1\ansi\ansicpg1252\deff0\nouicompat\deflang1033{\fonttbl{\f0\fnil\fcharset0 Calibri;}}
{\*\generator Riched20 10.0.18362}\viewkind4\uc1 
\pard\sa200\sl276\slmult1\f0\fs22\lang9 CREATE DATABASE `darts` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci */ /*!80016 DEFAULT ENCRYPTION='N' */;\par
\par
CREATE TABLE `player` (\par
  `name` varchar(45) NOT NULL,\par
  `team` varchar(45) NOT NULL,\par
  `phone` varchar(14) NOT NULL DEFAULT '(336) 777-7777',\par
  PRIMARY KEY (`name`)\par
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;\par
\par
CREATE TABLE `player_stats` (\par
  `ID` int NOT NULL AUTO_INCREMENT,\par
  `week` int NOT NULL,\par
  `team` varchar(45) NOT NULL,\par
  `player` varchar(45) NOT NULL,\par
  `five_singles_wldnp` varchar(4) NOT NULL DEFAULT 'DNP',\par
  `five_singles_win` int NOT NULL,\par
  `five_singles_loss` int NOT NULL,\par
  `cricket_singles_wldnp` varchar(4) NOT NULL DEFAULT 'DNP',\par
  `cricket_singles_win` int NOT NULL,\par
  `cricket_singles_loss` int NOT NULL,\par
  `five_doubles_wldnp` varchar(4) NOT NULL DEFAULT 'DNP',\par
  `five_doubles_win` int NOT NULL,\par
  `five_doubles_loss` int NOT NULL,\par
  `cricket_doubles_wldnp` varchar(4) NOT NULL DEFAULT 'DNP',\par
  `cricket_doubles_win` int NOT NULL,\par
  `cricket_doubles_loss` int NOT NULL,\par
  `plusone` int NOT NULL DEFAULT '0',\par
  `toneighty` int NOT NULL DEFAULT '0',\par
  `ninehtr` int NOT NULL DEFAULT '0',\par
  `mvp` int NOT NULL DEFAULT '0',\par
  `highout` int NOT NULL DEFAULT '0',\par
  `sixbulls` int NOT NULL DEFAULT '0',\par
  PRIMARY KEY (`ID`)\par
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='\tab\tab\tab\tab\tab\tab\tab\tab\tab\tab\tab\tab\tab\tab\tab\tab\tab\tab\tab\tab\tab\tab\tab\tab\tab\tab\tab\tab\tab\tab\tab\tab\tab\tab CREATE TABLE `schedule` (\par
  `week` int NOT NULL,\par
  `date` date NOT NULL,\par
  `venue1` varchar(45) NOT NULL,\par
  `team1` varchar(45) NOT NULL,\par
  `team2` varchar(45) NOT NULL,\par
  `venue2` varchar(45) NOT NULL,\par
  `team3` varchar(45) NOT NULL,\par
  `team4` varchar(45) NOT NULL,\par
  `venue3` varchar(45) NOT NULL,\par
  `team5` varchar(45) NOT NULL,\par
  `team6` varchar(45) NOT NULL,\par
  `venue4` varchar(45) NOT NULL,\par
  `nteam7` varchar(45) NOT NULL,\par
  `teambye` varchar(3) NOT NULL DEFAULT 'BYE',\par
  PRIMARY KEY (`week`)\par
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;\par
\par
CREATE TABLE `team` (\par
  `team` varchar(45) NOT NULL,\par
  `abbreviation` varchar(2) NOT NULL,\par
  PRIMARY KEY (`team`)\par
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;\par
\par
CREATE TABLE `users` (\par
  `username` varchar(25) NOT NULL,\par
  `password` binary(100) NOT NULL,\par
  `email` varchar(45) NOT NULL,\par
  `uteam` varchar(45) NOT NULL,\par
  PRIMARY KEY (`username`)\par
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;\par
\tab\tab\tab\tab\tab\tab\tab\tab\tab\tab\tab\tab\tab\tab\tab\tab\tab\tab\tab\tab\tab\tab\tab\tab\tab\tab\tab\tab\tab\tab\tab\tab\tab\tab\tab\tab\tab\tab\tab\tab\tab\tab\tab\tab\tab\tab\tab\tab\tab\tab\tab\tab\tab\tab\tab\tab\tab\tab\tab\tab\tab\tab\tab\tab\tab\tab\tab\tab\tab\tab\tab\tab\tab\tab\tab\tab\tab\tab\tab\tab\tab\tab\tab\tab\tab\tab\tab\tab\tab\tab\tab\tab\tab\tab\tab\tab\tab\tab\tab\tab\tab\tab\tab\tab\tab\tab\tab\tab\tab\tab\tab\tab\tab\tab\tab\tab\tab\tab\tab\tab\tab\tab\tab\tab\tab\tab\tab\tab\tab\tab\tab\tab\par
}
 /*
-- Query: SELECT * FROM darts.users
LIMIT 0, 1000

-- Date: 2020-08-16 14:51
*/
INSERT INTO `` (`username`,`password`,`email`,`uteam`) VALUES ('',?,'','');
INSERT INTO `` (`username`,`password`,`email`,`uteam`) VALUES ('hteeter56',?,'g@g.p','Alpha');
INSERT INTO `` (`username`,`password`,`email`,`uteam`) VALUES ('jteeter45',?,'j@gmail.com','Alpha');
INSERT INTO `` (`username`,`password`,`email`,`uteam`) VALUES ('kteeter42',?,'this@wart','Beta');
/*
-- Query: SELECT * FROM darts.team
LIMIT 0, 1000

-- Date: 2020-08-16 14:49
*/
INSERT INTO `` (`team`,`abbreviation`) VALUES ('None','NA');
INSERT INTO `` (`team`,`abbreviation`) VALUES ('TeamA','TA');
INSERT INTO `` (`team`,`abbreviation`) VALUES ('TeamB','TB');
INSERT INTO `` (`team`,`abbreviation`) VALUES ('TeamC','TC');
INSERT INTO `` (`team`,`abbreviation`) VALUES ('TeamD','TD');
INSERT INTO `` (`team`,`abbreviation`) VALUES ('TeamE','TE');
INSERT INTO `` (`team`,`abbreviation`) VALUES ('TeamF','TF');
INSERT INTO `` (`team`,`abbreviation`) VALUES ('TeamG','TG');

ID,week,team,player,five_singles_wldnp,five_singles_win,five_singles_loss,cricket_singles_wldnp,cricket_singles_win,cricket_singles_loss,five_doubles_wldnp,five_doubles_win,five_doubles_loss,cricket_doubles_wldnp,cricket_doubles_win,cricket_doubles_loss,plusone,toneighty,ninehtr,mvp,highout,sixbulls
1,1,TeamA,PlayerA1,Win,1,0,Win,1,0,Win,1,0,Win,1,0,0,0,0,0,545,1
2,1,TeamB,PlayerB1,Loss,0,1,Loss,0,1,Loss,0,1,Loss,0,1,1,1,1,1,0,2
3,2,TeamA,PlayerA1,DNP,0,0,DNP,0,0,Win,1,0,Loss,0,1,0,0,0,0,0,0
4,3,TeamB,PlayerB5,Win,1,0,Win,1,0,Win,1,0,Win,1,0,0,0,0,1,0,0
6,3,TeamA,PlayerA3,Win,1,0,Loss,0,1,Loss,0,1,Win,1,0,2,3,4,0,545,5
7,1,TeamB,PlayerB3,Win,1,0,Win,1,0,Win,1,0,Win,1,0,0,0,0,0,0,0
8,3,TeamB,PlayerB3,Win,1,0,Loss,0,1,DNP,0,0,DNP,0,0,0,0,0,0,0,0
10,3,TeamA,PlayerA5,Win,1,0,Win,1,0,Win,1,0,Win,1,0,0,0,0,0,0,0;

/*
-- Query: SELECT * FROM darts.player
LIMIT 0, 1000

-- Date: 2020-08-16 14:35
*/
INSERT INTO `` (`name`,`team`,`phone`) VALUES ('PlayerA1','TeamA','(336) 777-7777');
INSERT INTO `` (`name`,`team`,`phone`) VALUES ('PlayerA2','TeamA','(336) 777-7777');
INSERT INTO `` (`name`,`team`,`phone`) VALUES ('PlayerA3','TeamA','(336) 777-7777');
INSERT INTO `` (`name`,`team`,`phone`) VALUES ('PlayerA4','TeamA','(336) 777-7777');
INSERT INTO `` (`name`,`team`,`phone`) VALUES ('PlayerA5','TeamA','(336) 777-7777');
INSERT INTO `` (`name`,`team`,`phone`) VALUES ('PlayerB1','TeamB','(336) 777-7777');
INSERT INTO `` (`name`,`team`,`phone`) VALUES ('PlayerB2','TeamB','(336) 777-7777');
INSERT INTO `` (`name`,`team`,`phone`) VALUES ('PlayerB3','TeamB','(336) 777-7777');
INSERT INTO `` (`name`,`team`,`phone`) VALUES ('PlayerB4','TeamB','(336) 777-7777');
INSERT INTO `` (`name`,`team`,`phone`) VALUES ('PlayerB5','TeamB','(336) 777-7777');
INSERT INTO `` (`name`,`team`,`phone`) VALUES ('PlayerC1','TeamC','(336) 777-7777');
INSERT INTO `` (`name`,`team`,`phone`) VALUES ('PlayerC2','TeamC','(336) 777-7777');
INSERT INTO `` (`name`,`team`,`phone`) VALUES ('PlayerC3','TeamC','(336) 777-7777');
INSERT INTO `` (`name`,`team`,`phone`) VALUES ('PlayerC4','TeamC','(336) 777-7777');

week,date,venue1,team1,team2,venue2,team3,team4,venue3,team5,team6,venue4,nteam7,teambye
0,2020-01-01,None,None,None,None,None,None,None,None,None,None,None,Bye
1,2020-03-01,VenueA,TeamA,TeamB,VenueC,TeamC,TeamD,VenueE,TeamE,TeamF,VenueG,TeamG,Bye
2,2020-03-08,VenueB,TeamB,TeamC,VenueD,TeamD,TeamE,VenueF,TeamF,TeamG,VenueA,TeamA,Bye
3,2020-03-16,VenueC,TeamC,TeamD,VenueE,TeamE,TeamF,VenueG,TeamG,TeamA,VenueB,TeamB,Bye

END of SQL statements
 
