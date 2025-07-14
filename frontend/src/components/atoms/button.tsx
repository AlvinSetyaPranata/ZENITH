interface ButtonProps {
    title: string;
    onClick?: () => void;
}

export default function Button({ title, onClick } : ButtonProps) {
  return (
    <button on:click={onClick} class="rounded-md w-[300px] py-3 text-center bg-white text-black font-medium">{title}</button>
  )
}
