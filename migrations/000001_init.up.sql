CREATE OR REPLACE FUNCTION "update_modifiedAt"()
    RETURNS TRIGGER AS $$
BEGIN
    NEW."modifiedAt" = now();
    RETURN NEW;
END;
$$
LANGUAGE 'plpgsql';

CREATE SCHEMA "snowboards";

CREATE TABLE "snowboards"."persons"
(
    "id"         INTEGER PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY,
    "name"       TEXT                                   NOT NULL,

    "createdAt"  TIMESTAMP WITH TIME ZONE DEFAULT now() NOT NULL,
    "modifiedAt" TIMESTAMP WITH TIME ZONE DEFAULT now() NOT NULL
);

INSERT INTO "snowboards"."persons" ("id", "name")
VALUES
(  1, 'Jake Burton Carpenter'),
(  2, 'Jeremy Jones'),
(  3, 'Chris Prior'),
(  4, 'Bob Carlson'),
(  5, 'Abel Rossignol'),
(  6, 'Jason Brown'),
(  7, 'Blue Montgomery'),
(  8, 'Ken Block'),
(  9, 'Sean Genovese'),
( 10, 'Jeff Keenan'),
( 11, 'Max Jenke'),
( 12, 'Ben Dietermann'),
( 13, 'Jeff Baughn'),
( 14, 'Peter Bauer'),
( 15, 'Anian Thrainer'),
( 16, 'Selim Abdi'),
( 17, 'Klem Branner'),
( 18, 'Lisa Branner'),
( 19, 'Ben Hall'),
( 20, 'Tommy Delago'),
( 21, 'Sepp Ardelt'),
( 22, 'Mats Drougge'),
( 23, 'GiGi RÜF'),
( 24, 'Jorgen Karlson'),
( 25, 'Alister Horn'),
( 26, 'Steph Nitsch'),
( 27, 'Laurel Nelson'),
( 28, 'Régis Rolland')
       ;

-- https://en.wikipedia.org/wiki/List_of_ski_brands
-- https://www.liveabout.com/list-of-major-snowboard-brands-3020383
-- https://snowboardingprofiles.com/where-are-snowboards-made-the-major-brands
-- https://www.skiutah.com/blog/authors/yeti/five-snowboard-companies-you-didn-t
CREATE TABLE "snowboards"."brands"
(
    "id"           INTEGER PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY,
    "name"         TEXT                                   NOT NULL,
    "websiteUrl"   TEXT                                   NOT NULL,
    "founders"     INTEGER[] DEFAULT '{}' NOT NULL,
    "foundedIn"    INTEGER,
    "originsFrom"  TEXT,

    "createdAt"  TIMESTAMP WITH TIME ZONE DEFAULT now() NOT NULL,
    "modifiedAt" TIMESTAMP WITH TIME ZONE DEFAULT now() NOT NULL
);

INSERT INTO "snowboards"."brands" ("id", "name", "websiteUrl", "founders", "foundedIn", "originsFrom")
VALUES
    (  1, 'Academy', 'https://www.academysnowboards.com', '{13}', 2002, 'Tahoe, California, USA'),
    (  2, 'Amplid', 'https://www.amplid.com', '{14,15}', 2005, 'Munich, Germany'),
    (  3, 'Anticonf', 'https://www.anticonf.com', '{16}', 2007, 'Switzerland'),
    (  4, 'Arbor', 'https://www.arborcollective.com', '{4}', 1995, 'Venice, CA, United States'),

    ( 10, 'Bataleon', 'https://www.bataleon.com', '{24}', 2000, 'Netherlands'),
    ( 11, 'Borealis', 'https://www.borealis-snowboards.com', '{19}', 2013, 'France'),
    ( 12, 'Buddy Buddy', 'https://buddybuddy.eu', '{}', NULL, 'Germany'),
    ( 13, 'Burton', 'https://www.burton.com', '{1}', 1977, 'USA'),

    ( 20, 'Canary Cartel', 'https://canarycartel.com', '{}', 2018, 'Innsbruck, Austria'),
    ( 21, 'CAPiTA', 'https://capitasnowboarding.com', '{6,7}', 2000, 'USA'),
    ( 22, 'Chimera', 'https://chimerasnowboards.com', '{25}', 2010, 'Salt Lake City, UT, USA'),

    ( 30, 'DC', 'https://www.dcshoes.com/snow/', '{8}', 1994, 'USA'),

    ( 40, 'Deus Ex Machina', '','{}',  NULL, NULL),

    ( 50, 'Dinosaurs Will Die', 'https://dinosaurswilldie.com', '{9,10}', 2006, 'Canada'),
    ( 51, 'Drake', 'https://www.northwave.com', '{}', 1997, NULL),
    ( 52, 'Dupraz', '', '{}', NULL, NULL),

    ( 60, 'Easy', 'https://easyboardcompany.com', '{}', NULL, NULL),
    ( 61, 'Elevated Surfcraft', '', '{}', NULL, NULL),
    ( 62, 'Endeavor', 'https://www.endeavorsnowboards.com', '{11}', 2001, 'Vancouver, Canada'),

    ( 70, 'Fjell', '', '{}', NULL, NULL),
    ( 71, 'Flow', '', '{}', NULL, NULL),
    ( 72, 'Forum', '', '{}', NULL, NULL),
    ( 73, 'Furberg', '', '{}', NULL, NULL),

    ( 80, 'Gara', '', '{}', NULL, NULL),
    ( 81, 'Gentemstick', '', '{}', NULL, NULL),
    ( 82, 'Goodboards', '', '{}', NULL, NULL),
    ( 83, 'GNU', 'https://www.gnu.com', '{}', 1977, NULL),

    ( 89, 'HEAD', '', '{}', NULL, NULL),

    ( 90, 'Jones', 'https://www.jonessnowboards.com', '{2}', 2010, NULL),

    (100, 'K2', '', '{}', NULL, NULL),
    (101, 'Konvoi', 'https://www.konvoisnowboards.com', '{12}', NULL, 'Czech Republic'),
    (102, 'Korua Shapes', '', '{}', NULL, NULL),

    (110, 'Lib Tech', 'https://www.lib-tech.com', '{}', NULL, NULL),
    (111, 'Lobster', '', '{}', NULL, NULL),

    (120, 'Marhar Snowboards', '', '{}', NULL, NULL),
    (121, 'Monument Snowboards', '', '{}', 2001, NULL),
    (122, 'Moss Snowstick', '', '{}', NULL, 'Japan'),
    (123, 'Moonchild', '', '{}', NULL, NULL),

    (130, 'Never Summer', '', '{}', NULL, NULL),
    (131, 'Niche', '', '{}', NULL, NULL),
    (132, 'Nidecker', '', '{}', NULL, NULL),
    (133, 'Nitro', 'https://nitrousa.com', '{20,21}', 1990, 'Seattle, Washington, USA'),
    (134, 'Nobile', '', '{}', NULL, NULL),
    (135, 'Notice', 'http://www.noticesnowboards.com/', '{}', NULL, 'Whitefish, Montana, USA'),

    (137, 'Option', '', '{}', 1999, 'Vancouver, Canada'),

    (140, 'Pallas', 'http://www.pallassnowboards.com/', '{26}', 2013, 'Salt Lake City, UT, USA'),
    (141, 'Play', '', '{}', NULL, NULL),
    (142, 'Plum', '', '{}', NULL, NULL),
    (143, 'Pogo', '', '{}', NULL, NULL),
    (144, 'Powfinder', '', '{}', NULL, NULL),
    (145, 'Prior', 'https://www.priorsnow.com', '{3}', 1989, NULL),
    (146, 'Public Snowboards', '', '{}', NULL, NULL),

    (150, 'Rad Air', '', '{}', NULL, NULL),
    (151, 'RICE28', 'https://www.rice28jp.com', '{}', NULL, 'Japan'),
    (152, 'Ride', '', '{}', NULL, NULL),
    (153, 'Rhyme Backcountry', 'http://www.rhymebc.com', '{28}', 2015, 'France'),
    (154, 'Rome', '', '{}', NULL, NULL),
    (155, 'Rossignol', 'https://www.rossignol.com', '{5}', 1907, 'France'),
    (156, 'Roxy', '', '{}', NULL, NULL),

    (160, 'Salomon', '', '{}', NULL, NULL),
    (161, 'Sandy', '', '{}', NULL, NULL),
    (162, 'Santa Cruz', '', '{}', NULL, NULL),
    (163, 'Silbaerg', '', '{}', NULL, NULL),
    (164, 'Sims', '', '{}', NULL, NULL),
    (165, 'Slash', 'https://slashsnow.com', '{23}', 2012, 'Switzerland'),
    (166, 'Stone', '', '{}', NULL, NULL),
    (167, 'Stranda', 'https://www.strandasnowboards.com', '{22}', 2015, 'Sweden'),

    (170, 'Telos', '', '{}', NULL, NULL),
    (171, 'The Interior Plain Project', '', '{}', NULL, NULL),

    (180, 'United Shapes', '', '{}', NULL, NULL),
    (181, 'Unity', 'https://www.unitysnowboards.com', '{}', 1995, 'USA'),

    (190, 'Venture', 'https://venturesnowboards.com', '{17,18}', 1999, 'USA'),
    (191, 'Verdad', '', '{}', NULL, NULL),
    (192, 'Vimana', '', '{}', NULL, NULL),

    (195, 'West', '', '{}', NULL, NULL),
    (196, 'Weston', '', '{}', NULL, NULL),
    (197, 'Wired', '', '{}', NULL, NULL),

    (200, 'Yes.', 'https://www.yesnowboard.com', '{}', NULL, NULL),
    (201, 'Yunika', '', '{}', NULL, NULL);

CREATE TABLE "snowboards"."catalogues"
(
    "id"           INTEGER PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY,
    "brandId"      INTEGER NOT NULL,
    "season"       TEXT NOT NULL,
    "type"         TEXT NOT NULL,
    "url"          TEXT NOT NULL,
    "size"         INTEGER,
    "pages"        INTEGER,

    "createdAt"  TIMESTAMP WITH TIME ZONE DEFAULT now() NOT NULL,
    "modifiedAt" TIMESTAMP WITH TIME ZONE DEFAULT now() NOT NULL,

    CONSTRAINT "catalogues_brandId_fkey" FOREIGN KEY ("brandId") REFERENCES "snowboards"."brands" ON DELETE CASCADE
);

INSERT INTO "snowboards"."catalogues" ("id", "brandId", "season", "type", "url", "size", "pages")
VALUES
       (  1,  90, 'W1011', 'issuu', 'https://issuu.com/jonessnowboards/docs/jones-1011', 0, 12),
       (  2,  90, 'W1112', 'issuu', 'https://issuu.com/jonessnowboards/docs/jones-1112', 0, 16),
       (  3,  90, 'W1213', 'issuu', 'https://issuu.com/jonessnowboards/docs/jones-1213', 0, 24),
       (  4,  90, 'W1314', 'issuu', 'https://issuu.com/jonessnowboards/docs/winter-jones_014__1_', 0, 48),
       (  5,  90, 'W1415', 'issuu', 'https://issuu.com/jonessnowboards/docs/jones-1415', 0, 88),
       (  6,  90, 'W1516', 'issuu', 'https://issuu.com/jonessnowboards/docs/jones_catalog_2015-2016', 0, 84),
       (  7,  90, 'W1617', 'issuu', 'https://issuu.com/jonessnowboards/docs/jones-1617_final', 0, 88),
       (  8,  90, 'W1718', 'issuu', 'https://issuu.com/jonessnowboards/docs/jones_1718_v1', 0, 84),
       (  9,  90, 'W1819', 'issuu', 'https://issuu.com/jonessnowboards/docs/jones-1819__2_', 0, 108),
       ( 10,  90, 'W1920', 'issuu', 'https://issuu.com/jonessnowboards/docs/jones-catalogue-1920-web', 0, 124),
       ( 11,  90, 'W2021', 'issuu', 'https://issuu.com/zuzupopo/docs/jones-catalogue-2021-small', 0, 144)
       ;

CREATE TABLE "snowboards"."snowboards"
(
    "id"        INTEGER PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY,
    "brandId"   INTEGER NOT NULL,
    "name"      TEXT NOT NULL,
    "season"    TEXT NOT NULL,
    "type"      TEXT NOT NULL,

    "createdAt"  TIMESTAMP WITH TIME ZONE DEFAULT now() NOT NULL,
    "modifiedAt" TIMESTAMP WITH TIME ZONE DEFAULT now() NOT NULL,

    CONSTRAINT "snowboards_brandId_fkey" FOREIGN KEY ("brandId") REFERENCES "snowboards"."brands" ON DELETE CASCADE
);

INSERT INTO "snowboards"."snowboards" ("id", "brandId", "name", "season", "type")
    VALUES
    (  1,  90, 'Flagship', 'W1011', 'snowboard'),
    (  2,  90, 'Flagship', 'W1112', 'snowboard'),
    (  3,  90, 'Flagship', 'W1213', 'snowboard'),
    (  4,  90, 'Flagship', 'W1314', 'snowboard'),
    (  5,  90, 'Flagship', 'W1415', 'snowboard'),
    (  6,  90, 'Flagship', 'W1516', 'snowboard'),
    (  7,  90, 'Flagship', 'W1617', 'snowboard'),
    (  8,  90, 'Flagship', 'W1718', 'snowboard'),
    (  9,  90, 'Flagship', 'W1819', 'snowboard'),
    ( 10,  90, 'Flagship', 'W1920', 'snowboard'),
    ( 11,  90, '10th Anniversary Flagship', 'W1920', 'snowboard'),
    ( 12,  90, 'Flagship', 'W2021', 'snowboard'),
    ( 13,  90, 'Aviator', 'W1314', 'snowboard'),
    ( 14,  90, 'Aviator', 'W1415', 'snowboard'),
    ( 15,  90, 'Aviator', 'W1516', 'snowboard'),
    ( 16,  90, 'Aviator', 'W1617', 'snowboard'),
    ( 17,  90, 'Aviator', 'W1718', 'snowboard'),
    ( 18,  90, 'Aviator', 'W1819', 'snowboard'),
    ( 19,  90, 'Aviator', 'W1920', 'snowboard'),
    ( 20,  90, 'Aviator', 'W1920', 'snowboard'),
    ( 21,  90, 'Aviator', 'W2021', 'snowboard')
    ;

CREATE TABLE "snowboards"."images"
(
    "id"            INTEGER PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY,
    "snowboardId"   INTEGER NOT NULL,
    "url"           TEXT NOT NULL,
    "size"          TEXT,
    "colorOfBase"   TEXT,

    "createdAt"  TIMESTAMP WITH TIME ZONE DEFAULT now() NOT NULL,
    "modifiedAt" TIMESTAMP WITH TIME ZONE DEFAULT now() NOT NULL,

    CONSTRAINT "images_snowboardId_fkey" FOREIGN KEY ("snowboardId") REFERENCES "snowboards"."snowboards" ON DELETE CASCADE
);

INSERT INTO "snowboards"."images" ("snowboardId", "url", "size", "colorOfBase")
    VALUES
    (20, 'http://images.snowboards.zaibatsu.ru/Jones/2020/aviator-01.png', NULL, 'black'),
    (20, 'http://images.snowboards.zaibatsu.ru/Jones/2020/aviator-02.png', NULL, 'blue')
    ;