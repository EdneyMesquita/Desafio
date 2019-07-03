CREATE TABLE "users" (
	"id" serial NOT NULL,
	"uuiduser" uuid NOT NULL,
	"name" varchar(100) NOT NULL,
	"email" varchar(100) NOT NULL,
	"password" varchar(255) NOT NULL,
	"cpf" varchar(15) NOT NULL,
	"avatar" integer NOT NULL,
	"datastart" varchar NOT NULL,
	CONSTRAINT "users_pk" PRIMARY KEY ("id")
) WITH (
  OIDS=FALSE
);



CREATE TABLE "avatar" (
	"id" serial NOT NULL,
	"url" varchar(255) NOT NULL,
	"type" varchar(20) NOT NULL,
	CONSTRAINT "avatar_pk" PRIMARY KEY ("id")
) WITH (
  OIDS=FALSE
);



ALTER TABLE "users" ADD CONSTRAINT "users_fk0" FOREIGN KEY ("avatar") REFERENCES "avatar"("id");


