import { Box, Container, CssBaseline, Divider, Grid, IconButton, List, ListItemButton, ListItemIcon, ListItemText, Toolbar, Typography } from "@mui/material"
import { createTheme, ThemeProvider } from "@mui/material/styles"
import { useState } from "react"
import { Balance, BarChart, ChevronLeft, Dashboard, Logout, Menu, People, Wallet } from '@mui/icons-material'
import MyAppBar from '../components/MyAppBar'
import MyDrawer from '../components/MyDrawer'
import { useDispatch, useSelector } from "react-redux"
import Copyright from "../components/Copyright"
import { setAuth, setMenu, setName, setPosition, setRole } from "../store/positionSlice"
import { dashboardLogout } from "../api/be"

const theme = createTheme()

export default function DashboardLayout({ children }) {
    const dispatch = useDispatch()
    const name = useSelector(state => state.position.name), role = useSelector(state => state.position.role), menu = useSelector(s => s.position.menu), auth = useSelector(s => s.position.auth)
    const [open, setOpen] = useState(true), [loading, setLoading] = useState(false)
    const toggleDrawer = () => setOpen(!open)
    const loggingOut = () => {
        setLoading(true)
        dashboardLogout(auth).then(r => {
            dispatch(setPosition('login'))
            dispatch(setMenu(0))
            dispatch(setName(''))
            dispatch(setRole(''))
            dispatch(setAuth(''))
        }).catch(console.log).finally(() => setLoading(false))
    }
    return <ThemeProvider theme={theme}>
        <Box sx={{ display: 'flex' }}>
            <CssBaseline />
            <MyAppBar open={open} position="absolute">
                <Toolbar sx={{ pr: '24px' }}>
                    <IconButton edge="start" color="inherit" aria-label="open drawer" onClick={toggleDrawer} sx={{ marginRight: '36px', ...(open && { display: 'none' }), }}>
                        <Menu />
                    </IconButton>
                    <Typography component="h1" variant="h6" color="inherit" noWrap sx={{ flexGrow: 1 }}>{name}'s Dashboard</Typography>
                </Toolbar>
            </MyAppBar>
            <MyDrawer variant="permanent" open={open}>
                <Toolbar sx={{ display: 'flex', alignItems: 'center', justifyContent: 'flex-end', px: [1], }}>
                    <IconButton onClick={toggleDrawer}><ChevronLeft /></IconButton>
                </Toolbar>
                <Divider />
                <List component='nav'>
                    <ListItemButton disabled={loading} onClick={() => dispatch(setMenu(0))} selected={menu === 0}>
                        <ListItemIcon><BarChart /></ListItemIcon>
                        <ListItemText primary='Charts' />
                    </ListItemButton>
                    <ListItemButton disabled={loading} onClick={() => dispatch(setMenu(1))} selected={menu === 1}>
                        <ListItemIcon><Wallet /></ListItemIcon>
                        <ListItemText primary='Budgets' />
                    </ListItemButton>
                    <ListItemButton disabled={loading} onClick={() => dispatch(setMenu(2))} selected={menu === 2}>
                        <ListItemIcon><Dashboard /></ListItemIcon>
                        <ListItemText primary='Categories' />
                    </ListItemButton>
                    <ListItemButton disabled={loading} onClick={() => dispatch(setMenu(3))} selected={menu === 3}>
                        <ListItemIcon><Balance /></ListItemIcon>
                        <ListItemText primary='Expenses' />
                    </ListItemButton>
                    {
                        role === 'Admin' ?
                            <>
                                <Divider sx={{ my: 1 }} />
                                <ListItemButton disabled={loading} onClick={() => dispatch(setMenu(4))} selected={menu === 4}>
                                    <ListItemIcon><People /></ListItemIcon>
                                    <ListItemText primary='Users' />
                                </ListItemButton>
                            </>
                            : <></>
                    }
                    <Divider sx={{ my: 1 }} />
                    <ListItemButton disabled={loading} onClick={() => loggingOut()}>
                        <ListItemIcon><Logout /></ListItemIcon>
                        <ListItemText primary='Logout' />
                    </ListItemButton>
                </List>
            </MyDrawer>
            <Box sx={{ display: 'flex', alignItems: 'center', justifyContent: 'flex-end', px: [1], }}>
                <Toolbar />
                <Container maxWidth="lg" sx={{ mt: 4, mb: 4 }}>
                    <Grid container spacing={3}>
                        {children}
                    </Grid>
                    <Copyright sx={{ pt: 4 }} />
                </Container>
            </Box>
        </Box>
    </ThemeProvider>
}