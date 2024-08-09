import { Grid, Paper, Table, TableBody, TableCell, TableHead, TableRow, Typography } from "@mui/material"
import { useEffect, useState } from "react"
import { useSelector } from "react-redux"
import { adminDeleteUser, adminUsers } from "../api/be"
import Confirmation from "../components/Confirmation"

export default function Users() {
    const [datas, setDatas] = useState([]), [loading, setLoading] = useState(false)
    const auth = useSelector(s => s.position.auth)
    const deleting = id => {
        setLoading(true)
        adminDeleteUser(auth, id).then(console.log).catch(console.log).finally(reloading)
    }
    const reloading = () => {
        setLoading(true)
        adminUsers(auth).then(r => setDatas(r.data)).catch(console.log).finally(() => setLoading(false))
    }
    useEffect(() => {
        reloading()
    }, [])
    return <Grid item xs={12}>
        <Paper sx={{ p: 2, display: 'flex', flexDirection: 'column' }}>
            <Typography component="h2" variant="h6" color="primary" gutterBottom>Users</Typography>
            <Table size="small">
                <TableHead>
                    <TableRow>
                        <TableCell>No</TableCell>
                        <TableCell>Full Name</TableCell>
                        <TableCell>User Name</TableCell>
                        <TableCell>Email</TableCell>
                        <TableCell>Actions</TableCell>
                    </TableRow>
                </TableHead>
                <TableBody>
                    {datas.map((v, i) => (
                        <TableRow key={v.id}>
                            <TableCell>{i + 1}</TableCell>
                            <TableCell>{v.fullname}</TableCell>
                            <TableCell>{v.username}</TableCell>
                            <TableCell>{v.email}</TableCell>
                            <TableCell>
                                <Confirmation onYes={() => deleting(v.id)} onClose={() => reloading()} buttonText='Delete' />
                            </TableCell>
                        </TableRow>
                    ))}
                </TableBody>
            </Table>
        </Paper>
    </Grid>
}