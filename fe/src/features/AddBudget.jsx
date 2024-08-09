import { Add } from "@mui/icons-material"
import { Button, Dialog, DialogActions, DialogContent, DialogTitle, Fab, TextField } from "@mui/material"
import { useState } from "react"
import { addBudget } from "../api/be"
import { useSelector } from "react-redux"

export default function AddBudget({ onClose }) {
    const auth = useSelector(s => s.position.auth)
    const [open, setOpen] = useState(false), [name, setName] = useState(''), [amount, setAmount] = useState(0), [loading, setLoading] = useState(false)
    const closing = () => {
        setOpen(false)
        onClose()
    }
    const saving = () => {
        setLoading(true)
        addBudget(auth, name, parseFloat(amount)).then(r => closing()).catch(console.log).finally(() => setLoading(false))
    }
    const openning = () => setOpen(true)
    return <>
        <Fab onClick={() => openning()} color="primary"><Add /></Fab>
        <Dialog onClose={() => closing()} open={open}>
            <DialogTitle>Add Budget</DialogTitle>
            <DialogContent>
                <TextField disabled={loading} onChange={e => setName(e.target.value)} value={name} margin="normal" required fullWidth id="name" label="Name" name="name" autoComplete="name" autoFocus />
                <TextField disabled={loading} onChange={e => setAmount(e.target.value)} value={amount} type="number" margin="normal" required fullWidth id="amount" label="Amount" name="amount" autoComplete="amount" autoFocus />
            </DialogContent>
            <DialogActions>
                <Button disabled={loading} onClick={() => saving()}>SAVE</Button>
            </DialogActions>
        </Dialog>
    </>
}