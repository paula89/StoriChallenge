CREATE TABLE IF NOT EXISTS
        Transactions (
              Id            varchar(50) NOT NULL,
              UserId        varchar(50) NOT NULL,
              CreationDate  date        NOT NULL,
              Debit         double      NOT NULL,
              Credit        double      NOT NULL,
              PRIMARY KEY (Id)
       );

