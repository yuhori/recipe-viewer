import { DataGrid, GridRowsProp, GridColDef } from '@mui/x-data-grid';
import React from 'react';
import axios from 'axios';
import Styles from '../styles/Styles'
import { Button } from '@mui/material';
import BlenderIcon from '@mui/icons-material/Blender';



function List() {
  const [rows, setRows] = React.useState<GridRowsProp>([]);
  React.useEffect(() => {
    const url = 'http://ik1-432-48290.vs.sakura.ne.jp:8080/api/v1/recipes/list';
    axios.get(url).then((response) => {
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