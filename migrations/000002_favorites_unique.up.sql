ALTER TABLE dashboard.favorites
ADD CONSTRAINT favorites_session_id_name_key UNIQUE (session_id, name);
