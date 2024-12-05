INSERT INTO public.game_boxes (amount, created_at, updated_at) VALUES
	 (50000.00, NOW(), NOW()),
	 (500000.00, NOW(), NOW()),
	 (10000.00, NOW(), NOW()),
	 (100000.00, NOW(), NOW()),
	 (1000000.00, NOW(), NOW()),
	 (200000.00, NOW(), NOW()),
	 (20000.00, NOW(), NOW()),
	 (2000000.00, NOW(), NOW()),
	 (5000.00, NOW(), NOW());

INSERT INTO public.source_of_funds (source_name, created_at, updated_at) VALUES
	 ('BANK_TRANSFER', NOW(), NOW()),
	 ('CREDIT_CARD', NOW(), NOW()),
	 ('CASH', NOW(), NOW()),
	 ('REWARD', NOW(), NOW());

INSERT INTO public.transaction_types (type_name, created_at, updated_at) VALUES
	 ('TOP_UP', NOW(), NOW()),
	 ('TRANSFER', NOW(), NOW());
