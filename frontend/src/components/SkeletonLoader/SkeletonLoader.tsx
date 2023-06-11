import styles from './SkeletonLoader.module.css';

export default function SkeletonLoader() {
	return (
		<div className={styles.skeletonContainer}>
			<div className={`${styles.skeletonTeam} ${styles.skeleton}`}></div>
			<div className={`${styles.skeletonName} ${styles.skeleton}`}></div>
			<div className={`${styles.skeletonText} ${styles.skeleton}`}></div>
			<div className={`${styles.skeletonText} ${styles.skeleton}`}></div>
			<div className={`${styles.skeletonText} ${styles.skeleton}`}></div>
		</div>
	);
}
