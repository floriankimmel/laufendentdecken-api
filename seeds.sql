insert into reviews
    (
        id,
        product_name,
        brand,
        weight,
        price,
        podcast_episode,
        rating,
        statement
    )
values
    (
        "8ba03d9f-0cfc-4653-a1d5-0ff4e1ff02a5",
        "Evadict Trailunningschuh MT Cushion2",
        "Decathlon",
        350,
        119.99,
        "https://laufendentdecken-podcast.at/231/",
        1,
        "Gefällt uns"
    );

insert into review_shoes
    (
        id,
        review,
        drop,
        grip,
        sole
    )
values
    (
        "01966648-3376-4eea-9397-5089f3a07b19",
        "8ba03d9f-0cfc-4653-a1d5-0ff4e1ff02a5",
        4,
        "5",
        "Gummi"
    );

insert into review_pictures
    (
        id,
        review,
        link,
        alt_text
    )
values
    (
        "730c97f0-52fd-4be4-b998-66e46af8f64b",
        "8ba03d9f-0cfc-4653-a1d5-0ff4e1ff02a5",
        "https://laufendentdecken-podcast.at/wp-content/uploads/2023/11/IMG_1707-scaled.jpg",
        "Evadict Trailunningschuh MT Cushion2"
    );

insert into product_links
    (
        id,
        review,
        link,
        alt_text
    )
values
    (
        "05d3a474-bd86-4bcd-ade2-c8cffb052949",
        "8ba03d9f-0cfc-4653-a1d5-0ff4e1ff02a5",
        "https://www.decathlon.at/p/334419-196511-trailrunningschuhe-damen-mt-cushion-2-himbeerfarben.html",
        "Damen"
    );

insert into product_links
    (
        id,
        review,
        link,
        alt_text
    )
values
    (
        "0c1a20cd-0e47-4286-83d4-730e715d97de",
        "8ba03d9f-0cfc-4653-a1d5-0ff4e1ff02a5",
        "https://www.decathlon.at/p/334466-196688-trailrunningschuhe-herren-mt-cushion-2-turkis.html?queryID=80d2f23999a229d4281f5a30a50cfe79&objectID=4565196",
        "Herren"
    );

insert into trail_events
    (
        id,
        name,
        date,
        location,
        podcast_episode
    )
values
    (
        "7eecfeea-5070-42f7-ba1e-0536c8a55c53",
        "KUT80",
        "17.08.2024",
        "Sillian (Osttirol)",
        "https://laufendentdecken-podcast.at/220"
    );

insert into trail_event_distances
    (
        id,
        trail_event,
        distance,
        gpx_link
    )
values
    (
        "7abb382a-a9bc-466d-9e35-4839de9920d6",
        "7eecfeea-5070-42f7-ba1e-0536c8a55c53",
        80.0,
        "17.08.2024",
        "https://drive.google.com/file/d/1Aii5GAmFARHC1ibqI3HsIuRBtOHF5KTN/view?usp=sharing"
    );
