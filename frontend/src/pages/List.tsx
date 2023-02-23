import { DataGrid, GridRowsProp, GridColDef } from '@mui/x-data-grid';
import React from 'react';
import axios from 'axios';
import Styles from '../styles/Styles'
import { Button } from '@mui/material';
import BlenderIcon from '@mui/icons-material/Blender';


// const rows: GridRowsProp = [
//   {
//     id: "1",
//     name: '麻婆豆腐',
//     link: 'https://www.taishi-food.co.jp/recipes/item/15',
//   },
// ];


function List() {
  const [rows, setRows] = React.useState<GridRowsProp>([]);
  React.useEffect(() => {
    const url = 'http://localhost:8080/api/v1/recipes/list';
    axios.get(url).then((response) => {
      console.log(response.data);
      setRows(response.data);
    });
  }, []);

  // ボタンが押された時の処理
  const handleClick = (linkUrl: string) => {
    // リンク URL に移動
    window.location.href = linkUrl;
  };

  const columns: GridColDef[] = [
    // { field: 'id', headerName: 'id'},
    { field: 'name', headerName: '名前'},
    {
      field: 'openButton',
      headerName: '詳細',
      sortable: false,
      width: 90,
      renderCell: (params) => <Button variant="contained" color="primary" onClick={() => handleClick(params.row.link)} >詳細</Button>
    },
  ];

  return (
    <div className="App">
      <h2><BlenderIcon></BlenderIcon>Recipe Viewer<BlenderIcon></BlenderIcon></h2>
      <div style={{ height: 300, width: '100%' }}>
        <DataGrid
          getRowId={(row) => row.ID}
          sx={Styles.grid}
          rows={rows}
          columns={columns}
        />
      </div>
    </div>
  );
}
  
export default List;