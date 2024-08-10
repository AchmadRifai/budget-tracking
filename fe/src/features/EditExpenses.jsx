import { Edit } from "@mui/icons-material"
import { Button, Dialog, DialogActions, DialogContent, DialogTitle, Fab, FormControl, InputLabel, MenuItem, Select, TextField } from "@mui/material"
import { useEffect, useState } from "react"
import { editExpenses, preEditExpenses } from "../api/be"
import { useSelector } from "react-redux"
import { DateTimePicker } from "@mui/x-date-pickers"
import dayjs from "dayjs"

export default function EditExpenses({ dbData, onClose, disabled }) {
    const auth = useSelector(s => s.position.auth)
    const [open, setOpen] = useState(false), [loading, setLoading] = useState(false), [budgets, setBudgets] = useState([]), [categories, setCategories] = useState([])
    const [budget, setBudget] = useState(0), [category, setCategory] = useState(0), [amount, setAmount] = useState(dbData.amount), [time, setTime] = useState(dayjs.unix(dbData.time))
    const closing = () => {
        setOpen(false)
        onClose()
    }
    const saving = () => {
        setLoading(true)
        editExpenses(auth, category, budget, parseFloat(amount), time.unix(), dbData.id).then(r => closing()).catch(console.log).finally(() => setLoading(false))
    }
    const openning = () => setOpen(true)
    useEffect(() => {
        setLoading(true)
        preEditExpenses(auth).then(r => {
            setBudgets(r.budgets.data)
            setCategories(r.categories.data)
            setBudget(r.budgets.data.filter(v => v.name === dbData.budget).map(v => v.id)[0])
            setCategory(r.categories.data.filter(v => v.name === dbData.category).map(v => v.id)[0])
        }).catch(console.log).finally(() => setLoading(false))
    }, [])
    return <>
        <Fab disabled={disabled} onClick={() => openning()} color="secondary"><Edit /></Fab>
        <Dialog onClose={() => closing()} open={open}>
            <DialogTitle>Edit Category</DialogTitle>
            <DialogContent>
                <FormControl fullWidth disabled={loading} sx={{ p: 1 }}>
                    <InputLabel id='budget_lbl'>Budget</InputLabel>
                    <Select onChange={e => setBudget(e.target.value)} value={budget} label='Budget' labelId="budget_lbl">
                        {budgets.map(v => <MenuItem value={v.id}>{v.name}</MenuItem>)}
                    </Select>
                </FormControl>
                <FormControl fullWidth disabled={loading} sx={{ p: 1 }}>
                    <InputLabel id='category_lbl'>Category</InputLabel>
                    <Select onChange={e => setCategory(e.target.value)} value={category} label='Category' labelId="category_lbl">
                        {categories.map(v => <MenuItem value={v.id}>{v.name}</MenuItem>)}
                    </Select>
                </FormControl>
                <TextField disabled={loading} onChange={e => setAmount(e.target.value)} value={amount} type="number" margin="normal" required fullWidth id="amount" label="Amount" name="amount" autoComplete="amount" autoFocus />
                <DateTimePicker loading={loading} onChange={v => setTime(v)} value={time} label='Time' />
            </DialogContent>
            <DialogActions>
                <Button disabled={loading} onClick={() => saving()}>SAVE</Button>
            </DialogActions>
        </Dialog>
    </>
}