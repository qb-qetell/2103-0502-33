-- MySQL dump 10.17  Distrib 10.3.25-MariaDB, for debian-linux-gnu (x86_64)
--
-- Host: localhost    Database: 2103_0502_33
-- ------------------------------------------------------
-- Server version	10.3.25-MariaDB-0ubuntu0.20.04.1

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `car_model`
--

DROP TABLE IF EXISTS `car_model`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `car_model` (
  `record_id` char(8) DEFAULT NULL,
  `brand_name` varchar(32) DEFAULT NULL,
  `model_name` varchar(32) DEFAULT NULL,
  `id` char(5) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `car_model`
--

LOCK TABLES `car_model` WRITE;
/*!40000 ALTER TABLE `car_model` DISABLE KEYS */;
INSERT INTO `car_model` VALUES ('05222447','BMW','3 Series','00207'),('05223617','Cadillac','Cts','00228'),('05224341','Chevrolet','Aveo','00237'),('05224344','Chevrolet','Cruze','00241'),('05224346','Chevrolet','Equinox','00243'),('06004231','Honda','Accord','00332'),('06004234','Honda','Civic','00335'),('06004235','Honda','Cr-V','00336'),('06005752','Hyundai','Azera','00354'),('06005756','Hyundai','Elantra','00357'),('06005757','Hyundai','Genesis','00361'),('06005763','Hyundai','Santa Fe','00369'),('06005766','Hyundai','Sonata','00370'),('06005768','Hyundai','Tucson','00374'),('06143704','Kia','Cerato','00413'),('06143709','Kia','Optima','00419'),('06143714','Kia','Sorento','00428'),('06143717','Kia','Sportage','00431'),('06145039','Land Rover','Range Rover Evoque','00441'),('06145041','Land Rover','Range Rover Sport','00443'),('06160006','Mazda','Cx-9','00474'),('06160737','Mercedes-Benz','C300','07622'),('06160743','Mercedes-Benz','C250','00494'),('06160745','Mercedes-Benz','C300','00498'),('06163115','Mercedes-Benz','E200','00498'),('06164403','Mercedes-Benz','E300','00538'),('06164405','Mercedes-Benz','E350','00540'),('06164415','Mercedes-Benz','GL350','00557'),('06172801','Mitsubishi','Attrage','00620'),('06172805','Mitsubishi','L200','00633'),('06172810','Mitsubishi','Outlander','00637'),('06172811','Mitsubishi','Pajero','00638'),('06184201','Nissan','Almera','00642'),('06184202','Nissan','Altima','00643'),('06184207','Nissan','Frontier','00648'),('06184211','Nissan','Maxima','00651'),('06184220','Nissan','Rogue','00664'),('06184221','Nissan','Sentra','00665'),('06184228','Nissan','Versa','00673'),('06184229','Nissan','X-Trail','00674'),('06201510','Peugeot','607','00698'),('06203405','Renault','Logan','00714'),('06222402','Subaru','Outback','00742'),('06222603','Suzuki','BALENO','01864'),('06222608','Suzuki','Grand Vitara','00747'),('06224901','Toyota','4Runner','00759'),('06224907','Toyota','Camry','00765'),('06224911','Toyota','Corolla','00770'),('06224918','Toyota','Highlander','00778'),('06224919','Toyota','Hilux','00779'),('06224921','Toyota','Land Cruiser','00780'),('06224922','Toyota','Land Cruiser Prado','00781'),('06224928','Toyota','Rav4','00786'),('06224933','Toyota','Tundra','00791'),('05221813','Audi','A4','00185'),('05221844','Audi','A6','00187'),('05231835','Ford','Escape','00293'),('05231838','Ford','Explorer','00297'),('05231834','Ford','Edge','00292'),('05231843','Ford','Focus','00305'),('06142121','Jeep','Cherokee','00399'),('06142129','Jeep','Wrangler','00406'),('28035631','CHANGAN','CS95','07167'),('28035632','Citroen','C5','00268'),('28035633','Honda','Stream','00348'),('28035634','Hyundai','Eon','07164'),('28035635','Hyundai','Starex','07168'),('28035636','Jaguar','F Pace','07174'),('28035638','Kia','Picanto','00420'),('28035639','Mitsubishi','ASX','07175'),('28035641','Nissan','Bluebird','00645'),('28035642','Nissan','Kicks','07172'),('28035643','Nissan','Patrol','00659'),('28035644','Tata','Safari','06772'),('28035645','Toyota','SCION','00870'),('28035646','Toyota','Vitz','07157');
/*!40000 ALTER TABLE `car_model` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2021-03-28  4:55:50
