import React, { useState, useEffect, FC } from 'react';
import { makeStyles } from '@material-ui/core/styles';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import Paper from '@material-ui/core/Paper';
import Button from '@material-ui/core/Button';
import { DefaultApi } from '../../api/apis';
import { EntEatinghistory } from '../../api/models/EntEatinghistory'; // import interface Eatinghistory
import {
    Content,
    Header,
    Page,
   } from '@backstage/core';
import { Link as RouterLink } from 'react-router-dom';
import moment from 'moment';

 
const useStyles = makeStyles({
 table: {
   minWidth: 650,
 },
});

export default function EatinghistoryTable() {
 const classes = useStyles();
 const api = new DefaultApi();

 const [eatinghistorys, setEatinghistorys] = useState<EntEatinghistory[]>([]);


 const [loading, setLoading] = useState(true);

 useEffect(() => {
   const getEatinghistorys = async () => {
     const res = await api.listEatinghistory({ limit: 20, offset: 0 });
     setLoading(false);
     setEatinghistorys(res);
     console.log(res);
   };
   getEatinghistorys();
   
 }, [loading]);
 
 const deleteEatinghistorys = async (id: number) => {
   const res = await api.deleteEatinghistory({ id: id });
   setLoading(true);
 };
 
 return (
    <Page>
    <Header
    title="ประวัติการรับประทานอาหาร">
        <Button
                style={{ marginRight: 50 }}
                component={RouterLink}
                to="/"
                variant="contained"
              >
                BACK
             </Button>
    </Header>
    <Content>
   <TableContainer component={Paper}>
     <Table className={classes.table} aria-label="simple table">
       <TableHead>
         <TableRow>
           <TableCell align="center">ลำดับ</TableCell>
           <TableCell align="center">เวลาที่กิน</TableCell>
           <TableCell align="center">มื้ออาหาร</TableCell>
           <TableCell align="center">เมนูอาหาร</TableCell>
           <TableCell align="center">รสชาติ</TableCell>
           <TableCell align="center">ผู้ใช้</TableCell>
           <TableCell align="center">จัดการ</TableCell>
         </TableRow>
       </TableHead>
       <TableBody>
         {eatinghistorys.map((item:any) => (
           <TableRow key={item.id}>
             <TableCell align="center">{item.id}</TableCell>
             <TableCell align="center">{moment(item.addedTime).format('DD/MM/YYYY HH.mm  ')}</TableCell>
             <TableCell align="center">{item.edges?.mealplan?.id}</TableCell>
             <TableCell align="center">{item.edges?.foodmenu?.name}</TableCell>
             <TableCell align="center">{item.edges?.taste?.taste}</TableCell>
             <TableCell align="center">{item.edges?.user?.email}</TableCell>
             <TableCell align="center">
               <Button
                 onClick={() => {
                   deleteEatinghistorys(item.id);
                 }}
                 style={{ marginLeft: 10 }}
                 variant="contained"
                 color="secondary"
               >
                 Delete
               </Button>
             </TableCell>
           </TableRow>
         ))}
       </TableBody>
     </Table>
   </TableContainer>
   </Content>
   </Page>
 );
}