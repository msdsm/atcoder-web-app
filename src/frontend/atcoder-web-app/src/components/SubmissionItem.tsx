import { FC, memo } from 'react'
import { Submission } from '../types'

const SubmissionItemMemo: FC<Submission> = ({
    atcoder_id,
    time,
    problem,
    diff,
}) => {
    return (
        <tr>
            <td>{atcoder_id}</td>
            <td>{time}</td>
            <td>{problem}</td>
            <td>{diff}</td>
        </tr>
    )
}
export const SubmissionItem = memo(SubmissionItemMemo)