import LockOutlinedIcon from '@mui/icons-material/LockOutlined'
import { Avatar, Box, Button, Container, CssBaseline, Grid, Link, TextField, Typography } from '@mui/material'
import { createTheme, ThemeProvider } from '@mui/material/styles'
import Copyright from '../components/Copyright'
import { useDispatch } from 'react-redux'
import { setPosition } from '../store/positionSlice'

const defaultTheme = createTheme()

export default function Login() {
    const dispatch = useDispatch()
    return <ThemeProvider theme={defaultTheme}>
        <Container component='main' maxWidth='xs'>
            <CssBaseline />
            <Box sx={{ marginTop: 8, display: 'flex', flexDirection: 'column', alignItems: 'center', }}>
                <Avatar sx={{ m: 1, bgcolor: 'secondary.main' }}>
                    <LockOutlinedIcon />
                </Avatar>
                <Typography component='h1' variant='h5'>Sign In</Typography>
                <Box component='form' noValidate sx={{ mt: 1 }}>
                    <TextField margin="normal" required fullWidth id="email" label="Email Address" name="email" autoComplete="email" autoFocus />
                    <TextField margin="normal" required fullWidth name="password" label="Password" type="password" id="password" autoComplete="current-password" />
                    <Button type="submit" fullWidth variant="contained" sx={{ mt: 3, mb: 2 }}>Sign In</Button>
                    <Grid container>
                        <Grid item>
                            <Link onClick={() => dispatch(setPosition('register'))} href='#' variant='body2'>Sign Up</Link>
                        </Grid>
                    </Grid>
                </Box>
            </Box>
            <Copyright sx={{ mt: 8, mb: 4 }} />
        </Container>
    </ThemeProvider>
}