CREATE TABLE IF NOT EXISTS tbltodos(
      id                INT AUTO_INCREMENT NOT NULL,
      strtitle          VARCHAR(128) NOT NULL,
      strDescription    VARCHAR(128) NOT NULL,
      priority          INT NOT NULL,
      dtCompleted       DATE,  
      CreatedAt         DATE,  
      DueDate           DATE,
      PRIMARY KEY (`id`)
)