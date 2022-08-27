CREATE KEYSPACE oauth WITH replication = {'class':'SimpleStrategy', 'replication_factor' : 1};
/
CREATE TABLE oauth.user_auth ( ID VARCHAR , Nonce VARCHAR, Fingerprint VARCHAR, Created NUMBER, PRIMARY KEY( ID ) )
/