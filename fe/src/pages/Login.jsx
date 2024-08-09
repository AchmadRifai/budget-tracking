import LockOutlinedIcon from '@mui/icons-material/LockOutlined'
import { Avatar, Box, Button, Container, CssBaseline, Grid, Link, TextField, Typography } from '@mui/material'
import { createTheme, ThemeProvider } from '@mui/material/styles'
import Copyright from '../components/Copyright'
import { useDispatch } from 'react-redux'
import { setAuth, setPosition } from '../store/positionSlice'
import { useState } from 'react'
import { login } from '../api/be'

const defaultTheme = createTheme()

export default function Login() {
    const dispatch = useDispatch()
    const [loading, setLoading] = useState(false)
    const submit = e => {
        e.preventDefault()
        setLoading(true)
        const data = new FormData(e.currentTarget)
        login(data.get('username'), data.get('password')).then(r => {
            console.log(r)
            localStorage.setItem('auth', r.header.authorization)
            dispatch(setPosition('dashboard'))
            dispatch(setAuth(r.header.authorization))
        }).catch(e => {
            console.log(e);
        }).finally(() => {
            setLoading(false)
        })
    }
    return <ThemeProvider theme={defaultTheme}>
        <Container component='main' maxWidth='xs'>
            <CssBaseline />
            <Box sx={{ marginTop: 8, display: 'flex', flexDirection: 'column', alignItems: 'center', }}>
                <Avatar sx={{ m: 1, bgcolor: 'secondary.main' }}>
                    <LockOutlinedIcon />
                </Avatar>
                <Typography component='h1' variant='h5'>Sign In</Typography>
                <Box onSubmit={submit} component='form' noValidate sx={{ mt: 1 }}>
                    <TextField disabled={loading} margin="normal" required fullWidth id="username" label="User Name" name="username" autoComplete="username" autoFocus />
                    <TextField disabled={loading} margin="normal" required fullWidth name="password" label="Password" type="password" id="password" autoComplete="current-password" />
                    <Button disabled={loading} type="submit" fullWidth variant="contained" sx={{ mt: 3, mb: 2 }}>Sign In</Button>
                    <Grid container>
                        <Grid item>
                            <Link visibility={!loading} onClick={() => dispatch(setPosition('register'))} href='#' variant='body2'>Sign Up</Link>
                        </Grid>
                    </Grid>
                </Box>
            </Box>
            <Copyright sx={{ mt: 8, mb: 4 }} />
        </Container>
    </ThemeProvider>
}