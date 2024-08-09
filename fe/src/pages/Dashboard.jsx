import { Grid } from "@mui/material"
import DashboardLayout from "../layouts/Dashboard"

export default function Dashboard() {
    return <DashboardLayout>
        <Grid item xs={12} md={8} lg={9}></Grid>
    </DashboardLayout>
}