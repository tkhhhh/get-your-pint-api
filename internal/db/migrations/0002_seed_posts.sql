INSERT INTO posts (title, content, excerpt) VALUES
    ('First Blog Post',  'This is the content of the first blog post.',  'This is the first blog post.'),
    ('Second Blog Post', 'This is the content of the second blog post.', 'This is the second blog post.'),
    ('Third Blog Post',  'This is the content of the third blog post.',  'This is the third blog post.')
ON CONFLICT DO NOTHING;
