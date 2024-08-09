import { Grid, Paper, Typography } from "@mui/material"
import { PieChart } from '@mui/x-charts'

export default function MyPieChart({ title, data }) {
    return <Grid item xs={12}>
        <Paper sx={{ p: 2, display: 'flex', flexDirection: 'column' }}>
            <Typography component="h2" variant="h6" color="primary" gutterBottom>{title}</Typography>
            <PieChart series={[
                { data: Object.keys(data).map((k, i) => ({ label: k, id: i, value: Object.values(data[k]).reduce((a, n) => a + n, 0) })) }
            ]} width={400} height={200} />
        </Paper>
    </Grid>
}