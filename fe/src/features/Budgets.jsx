import { Grid, Paper, Table, TableBody, TableCell, TableHead, TableRow, Typography } from "@mui/material"
import { useEffect, useState } from "react"
import { budget, delBudget } from "../api/be"
import { useSelector } from "react-redux"
import AddBudget from "./AddBudget"
import EditBudget from "./EditBudget"
import Confirmation from "../components/Confirmation"

export default function Budgets() {
    const auth = useSelector(s => s.position.auth)
    const [datas, setDatas] = useState([]), [loading, setLoading] = useState(false)
    const reloading = () => {
        setLoading(true)
        budget(auth).then(r => setDatas(r.data)).catch(console.log).finally(() => setLoading(false))
    }
    const deleting = id => {
        setLoading(true)
        delBudget(auth, id).then(r => reloading()).catch(console.log).finally(() => setLoading(false))
    }
    useEffect(() => {
        reloading()
    }, [])
    return <Grid item xs={12}>
        <Paper sx={{ p: 2, display: 'flex', flexDirection: 'column' }}>
            <Typography component="h2" variant="h6" color="primary" gutterBottom>Budgets</Typography>
            <Table size="small">
                <TableHead>
                    <TableRow>
                        <TableCell>No</TableCell>
                        <TableCell>Name</TableCell>
                        <TableCell>Amount</TableCell>
                        <TableCell>Actions</TableCell>
                    </TableRow>
                </TableHead>
                <TableBody>
                    {
                        datas.map((v, i) => <TableRow>
                            <TableCell>{i + 1}</TableCell>
                            <TableCell>{v.name}</TableCell>
                            <TableCell>{v.amount}</TableCell>
                            <TableCell>
                                <EditBudget onClose={() => reloading()} dbData={v} disabled={loading} />
                                <Confirmation onYes={() => deleting(v.id)} buttonText='Delete' onClose={() => reloading()} disabled={loading} />
                            </TableCell>
                        </TableRow>)
                    }
                </TableBody>
            </Table>
            <AddBudget onClose={() => reloading()} />
        </Paper>
    </Grid>
}