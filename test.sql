-- ============================================================================== review product
--  ALL REVIEW
SELECT 
    p.name,
    i.url AS image,
    p.description,
    p.price,
    COUNT(rp.id) total_review
    FROM review_product rp
    JOIN transaction_details td ON rp.id_transaction_details = td.id
    JOIN products p ON td.product_id = p.id
    LEFT JOIN product_images pi ON p.id = pi.product_id
    LEFT JOIN images i ON pi.image_id = i.id
    GROUP BY p.id,p.name,i.url,p.description,p.price
    ORDER BY total_review DESC;

-- REVIEW BY ID 
SELECT 
    p.name,
    i.url AS image,
    p.description,
    p.price,
    COUNT(rp.id) total_review
    FROM review_product rp
    JOIN transaction_details td ON rp.id_transaction_details = td.id
    JOIN products p ON td.product_id = p.id
    LEFT JOIN product_images pi ON p.id = pi.product_id
    LEFT JOIN images i ON pi.image_id = i.id
	WHERE p.id=1
    GROUP BY p.id,p.name,i.url,p.description,p.price;


-- ================================================================================ recommeded product

SELECT 
    p.id,
    p.name,
    i.url,
    p.description,
    p.price,
    COUNT(rp.id) total_review,
    AVG(rp.rating) avg_rating
    FROM review_product rp
    JOIN transaction_details td ON rp.id_transaction_details = td.id
    JOIN products p ON td.product_id = p.id
    LEFT JOIN product_images pi ON p.id = pi.product_id
    LEFT JOIN images i ON pi.image_id = i.id
    GROUP BY p.id,p.name,i.url,p.description,p.price
    ORDER BY avg_rating DESC
    LIMIT 5;


SELECT 
    p.id,
    p.name,
    i.url AS image,
    p.description,
    p.price,
    COUNT(rp.id) AS total_review,
    AVG(rp.rating) AS avg_rating
    FROM review_product rp
    JOIN transaction_details td ON rp.id_transaction_details = td.id
    JOIN products p ON td.product_id = p.id
    LEFT JOIN product_images pi ON p.id = pi.product_id
    LEFT JOIN images i ON pi.image_id = i.id
    WHERE id=1
    GROUP BY p.id, p.name, i.url, p.description, p.price;