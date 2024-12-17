INSERT INTO apps (id, name, secret)
VALUES (1, 'test', 'secret-sample')
ON CONFLICT DO NOTHING;