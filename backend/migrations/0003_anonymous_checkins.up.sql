-- Anonymous check-ins until auth lands (signin page is still prototype).
ALTER TABLE checkins ALTER COLUMN user_id DROP NOT NULL;
