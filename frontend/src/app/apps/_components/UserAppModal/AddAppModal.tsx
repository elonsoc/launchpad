import styles from './AddAppModal.module.css';
import { useState } from 'react';

interface ModalProps {
	onAdd: (name: string, description: string, owners: string) => void;
	onClose: (isClosed: boolean) => void;
}

const AddAppModal = ({ onAdd, onClose }: ModalProps) => {
	const [state, setState] = useState({ name: '', description: '', owners: '' });

	const handleInputChange = (event: any) => {
		const { name, value } = event.target;
		setState((prevInfo) => ({
			...prevInfo,
			[name]: value,
		}));
	};

	return (
		<div className={styles.fullScreenContainer}>
			<div className={styles.modalWindow}>
				<form>
					<div className={styles.inputWrapper}>
						<label htmlFor='name'>Name</label>
						<input
							type='text'
							id='name'
							name='name'
							value={state.name}
							onChange={handleInputChange}
						></input>
					</div>
					<div className={styles.inputWrapper}>
						<label htmlFor='description'>Description</label>
						<input
							type='text'
							id='description'
							name='description'
							value={state.description}
							onChange={handleInputChange}
						></input>
					</div>
					<div className={styles.inputWrapper}>
						<label htmlFor='owners'>Owners</label>
						<input
							type='text'
							id='owners'
							name='owners'
							value={state.owners}
							onChange={handleInputChange}
						></input>
					</div>
					<button
						className={styles.addButton}
						type='submit'
						onClick={() => onAdd(state.name, state.description, state.owners)}
					>
						Add
					</button>
				</form>
				<button
					className={styles.closeButton}
					type='button'
					onClick={() => onClose(false)}
				>
					<svg xmlns='http://www.w3.org/2000/svg' viewBox='0 0 24 24'>
						<title>close</title>
						<path d='M19,6.41L17.59,5L12,10.59L6.41,5L5,6.41L10.59,12L5,17.59L6.41,19L12,13.41L17.59,19L19,17.59L13.41,12L19,6.41Z' />
					</svg>
				</button>
			</div>
		</div>
	);
};

export default AddAppModal;
