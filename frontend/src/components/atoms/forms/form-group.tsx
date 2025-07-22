import { ParentProps } from 'solid-js'


interface FormGroupPropsType extends ParentProps {
    title: string
}

export default function FormGroup(props: FormGroupPropsType) {
  return (
    <div class='w-full'>

        <h2 class='text-xl font-semibold text-white'>{props.title}</h2>

        <div class='grid'>
            {props.children}
        </div>
    </div>
  )
}
