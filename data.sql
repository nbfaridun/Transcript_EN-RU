SELECT * FROM sessions;

select * from sessions
limit 5 offset 10;

CREATE OR REPLACE FUNCTION findTotalRows()
RETURNS integer
    LANGUAGE plpgsql
    AS $$
    DECLARE
        total_num integer := 0;
        row_item record;
        my_cursor CURSOR FOR SELECT (id * merchant_id) as m FROM sessions
        offset 0 fetch next 5 rows only;
    BEGIN
        open my_cursor;
        loop
            fetch my_cursor into row_item;
            exit when not found;
            total_num := total_num + row_item.m;
        end loop;
        close my_cursor;
        return total_num;
    END;
$$;

CREATE OR REPLACE FUNCTION availabilityCheck()
    RETURNS text
    LANGUAGE plpgsql
AS $$
DECLARE
    total_num integer := 0;
    answer varchar;
    row_item record;
    my_cursor CURSOR FOR SELECT (id * merchant_id) as m FROM sessions
                         offset 0 fetch next 5 rows only;
BEGIN
    open my_cursor;
    loop
        fetch my_cursor into row_item;
        exit when not found;
        total_num := total_num + row_item.m;
    end loop;
    close my_cursor;
    if total_num < 4000 then
            return 'Available'; else return 'Unavailable';
        end if;
END;
$$;


drop function findTotalRows();


select findTotalRows();

select availabilityCheck();

