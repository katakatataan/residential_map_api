SELECT id, built_count, total_square_meter, year, month, residential_use_type_id, construction_type_id, city_id, build_type_id, residential_type_id, structure_type_id, pref_id, city_name, pref_name, to_char(build_date,'YYYY-MM') as build_date,
rut.name as residential_use_type,
ct.name as construction_type,
bt.name as build_type,
rt.name as residential_type,
st.name as structure_type
FROM city_data
LEFT JOIN
mst_residential_use_type as rut
ON
rut.id = city_data.residential_use_type
LEFT JOIN
mst_construction_type as ct
ON
ct.id = city_data.construction_type_id
LEFT JOIN
mst_build_type as bt
ON
bt.id = city_data.build_type_id
LEFT JOIN
mst_residential_type as rt
ON
rt.id = city_data.residential_type_id
LEFT JOIN
mst_structure_type as st
ON
st.id = city_data.structure_type_id
ORDER BY city_id ASC, build_date ASC