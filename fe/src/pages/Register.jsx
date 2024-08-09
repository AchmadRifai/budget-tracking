import { useDispatch } from "react-redux"
import { setAuth, setPosition } from "../store/positionSlice"
import { createTheme, ThemeProvider } from "@mui/material/styles"
import { Avatar, Box, Button, Container, CssBaseline, Grid, Link, TextField, Typography } from "@mui/material"
import Copyright from "../components/Copyright"
import LockOutlined from "@mui/icons-material/LockOutlined"
import { useState } from "react"
import { register } from "../api/be"

const theme = createTheme()

export default function Register() {
    const dispatch = useDispatch()
    const [loading, setLoading] = useState(false)
    const submit = e => {
        e.preventDefault()
        setLoading(true)
        const data = new FormData(e.currentTarget)
        register(`${data.get('firstName')} ${data.get('lastName')}`, data.get('username'), data.get('password'), data.get('email')).then(r => {
            console.log(r)
            localStorage.setItem('auth', r.header.authorization)
            dispatch(setPosition('dashboard'))
            dispatch(setAuth(r.header.authorization))
        }).catch(console.log).finally(() => setLoading(false))
    }
    return <ThemeProvider theme={theme}>
        <Container component="main" maxWidth="xs">
            <CssBaseline />
            <Box sx={{ marginTop: 8, display: 'flex', flexDirection: 'column', alignItems: 'center', }}>
                <Avatar sx={{ m: 1, bgcolor: 'secondary.main' }}><LockOutlined /></Avatar>
                <Typography component="h1" variant="h5">Registration</Typography>
                <Box onSubmit={submit} component='form' noValidate sx={{ mt: 3 }}>
                    <Grid container spacing={2}>
                        <Grid item xs={12} sm={6}>
                            <TextField autoComplete="given-name" name="firstName" disabled={loading} required fullWidth id="firstName" label="First Name" autoFocus />
                        </Grid>
                        <Grid item xs={12} sm={6}>
                            <TextField required fullWidth id="lastName" disabled={loading} label="Last Name" name="lastName" autoComplete="family-name" />
                        </Grid>
                        <Grid item xs={12}>
                            <TextField required fullWidth disabled={loading} id="username" label="Username" name="username" autoComplete="username" />
                        </Grid>
                        <Grid item xs={12}>
                            <TextField disabled={loading} required fullWidth id="email" label="Email Address" name="email" autoComplete="email" />
                        </Grid>
                        <Grid item xs={12}>
                            <TextField disabled={loading} required fullWidth name="password" label="Password" type="password" id="password" autoComplete="new-password" />
                        </Grid>
                    </Grid>
                    <Button disabled={loading} type='submit' fullWidth variant='contained' sx={{ mt: 3, mb: 2 }}>Register</Button>
                    <Grid container justifyContent='flex-end'>
                        <Link visibility={loading} onClick={() => dispatch(setPosition('login'))} href='#' variant='body2'>Log In</Link>
                    </Grid>
                </Box>
            </Box>
            <Copyright sx={{ mt: 5 }} />
        </Container>
    </ThemeProvider>
}