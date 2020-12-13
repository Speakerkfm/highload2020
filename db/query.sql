select u.id, u.username, u.bill, u.last_login_at, e.email, p.phone from "user" u
    left join email e on u.id = e.user_id
    left join phone p on u.id = p.user_id
    where (e.is_active=1 or e.is_active is null) and (p.is_active=1 or p.is_active is null);

with selling_items as (
    select i.name, l.id, l.buy_price, l.max_bet_price
    from lot l
        left join lot_item li on l.id = li.lot_id
        left outer join item i on i.id = li.item_id
    where i.data ->> 'type' = 'bow'
    ORDER BY l.buy_price DESC, l.max_bet_price DESC)
select u.username, b.total_price, b.created_at
from bet b
    left join selling_items si on si.id = b.lot_id
    left join "user" u on u.id = b.user_id
order by b.created_at;

select u.username, it.name, ii.count
from inventory_item ii
    left join item it on it.id = ii.item_id
    left join inventory i on ii.inventory_id=i.id
    left outer join game g on g.id = it.game_id
    left join "user" u on u.id = i.user_id
where u.username=(
    select u2.username
    from "user" u2
    order by u2.bill desc
    limit 1
) and g.name='World of Warcraft';

select u.username, b.lot_id, b.total_price, avg(b.total_price) OVER (PARTITION BY u.username)
from "user" u
left join bet b on u.id = b.user_id;