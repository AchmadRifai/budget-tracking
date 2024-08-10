import { Grid, Paper, Table, TableBody, TableCell, TableHead, TableRow, Typography } from "@mui/material"
import { useEffect, useState } from "react"
import { delExpenses, expenses } from "../api/be"
import { useSelector } from "react-redux"
import dayjs from 'dayjs'
import AddExpenses from "./AddExpenses"
import Convirmation from '../components/Confirmation'
import EditExpenses from "./EditExpenses"

export default function Expenses() {
    const auth = useSelector(s => s.position.auth)
    const [datas, setDatas] = useState([]), [loading, setLoading] = useState(false)
    const reloading = () => {
        setLoading(true)
        expenses(auth).then(r => setDatas(r.data)).catch(console.log).finally(() => setLoading(false))
    }
    const deleting = id => {
        setLoading(true)
        delExpenses(auth, id).then(r => reloading()).catch(console.log).finally(() => setLoading(false))
    }
    useEffect(() => {
        reloading()
    }, [])
    return <Grid item xs={12}>
        <Paper sx={{ p: 2, display: 'flex', flexDirection: 'column' }}>
            <Typography component="h2" variant="h6" color="primary" gutterBottom>Expenses</Typography>
            <Table size="small">
                <TableHead>
                    <TableRow>
                        <TableCell>No</TableCell>
                        <TableCell>Budget</TableCell>
                        <TableCell>Category</TableCell>
                        <TableCell>Amount</TableCell>
                        <TableCell>Time</TableCell>
                        <TableCell>Actions</TableCell>
                    </TableRow>
                </TableHead>
                <TableBody>
                    {
                        datas.map((v, i) => <TableRow>
                            <TableCell>{i + 1}</TableCell>
                            <TableCell>{v.budget}</TableCell>
                            <TableCell>{v.category}</TableCell>
                            <TableCell>{v.amount}</TableCell>
                            <TableCell>{dayjs.unix(v.time).toISOString()}</TableCell>
                            <TableCell>
                                <EditExpenses dbData={v} disabled={loading} onClose={() => reloading()} />
                                <Convirmation onYes={() => deleting(v.id)} buttonText='Delete' onClose={() => reloading()} disabled={loading} />
                            </TableCell>
                        </TableRow>)
                    }
                </TableBody>
            </Table>
            <AddExpenses onClose={() => reloading()} />
        </Paper>
    </Grid>
}