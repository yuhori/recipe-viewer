import { DataGrid, GridRowsProp, GridColDef } from '@mui/x-data-grid';
import PaidIcon from '@mui/icons-material/Paid';
import React from 'react';
import axios from 'axios';
import Styles from '../styles/Styles'


const rows: GridRowsProp = [
  {
    id: "1",
    title: '麻婆豆腐',
    link: 'https://www.taishi-food.co.jp/recipes/item/15',
  },
];

const columns: GridColDef[] = [
  { field: 'id', headerName: 'id', width: 150 },
  { field: 'title', headerName: 'title', width: 150 },
  { field: 'link', headerName: 'link', width: 150 },
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