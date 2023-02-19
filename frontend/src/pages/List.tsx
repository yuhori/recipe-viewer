import { DataGrid, GridRowsProp, GridColDef } from '@mui/x-data-grid';
import PaidIcon from '@mui/icons-material/Paid';
import React from 'react';
import axios from 'axios';
import Styles from '../styles/Styles'
import { Button } from '@mui/material';


const rows: GridRowsProp = [
  {
    id: "1",
    name: '麻婆豆腐',
    link: 'https://www.taishi-food.co.jp/recipes/item/15',
  },
];


function List() {
//   const [rows, setRows] = React.useState<GridRowsProp>([]);
//   React.useEffect(() => {
//     // cors の問題で、リクエストに失敗する
//     const url = 'https://emaxis.jp/web/api/v1.php?col=doc_pdf_all';
//     axios.get(url).then((response) => {
//       console.log(response.data);
//       setRows(response.data);
//     });
//   }, []);
  // ボタンが押された時の処理
  const handleClick = (linkUrl: string) => {
    // リンク URL に移動
    window.location.href = linkUrl;
  };

  const columns: GridColDef[] = [
    { field: 'id', headerName: 'id'},
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
      <h2><PaidIcon></PaidIcon>Recipe Viewer<PaidIcon></PaidIcon></h2>
      <div style={{ height: 300, width: '100%' }}>
        <DataGrid
        //   getRowId={(row) => row.fund_id}
          sx={Styles.grid}
          rows={rows}
          columns={columns}
        />
      </div>
    </div>
  );
}
  
export default List;