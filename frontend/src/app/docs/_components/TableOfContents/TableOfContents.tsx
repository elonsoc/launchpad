'use client';

import React, {
	useState,
	useRef,
	useEffect,
	SetStateAction,
	Dispatch,
} from 'react';

import styles from './TableOfContents.module.css';

const useIntersectionObserver = (
	setActiveId: Dispatch<SetStateAction<string | undefined>>
) => {
	const headingElementsRef = useRef<{
		[key: string]: IntersectionObserverEntry;
	}>({});
	useEffect(() => {
		const callback: IntersectionObserverCallback = (
			headings: IntersectionObserverEntry[]
		) => {
			headingElementsRef.current = headings.reduce((map, headingElement) => {
				map[headingElement.target.id] = headingElement;
				return map;
			}, headingElementsRef.current);

			// Get all headings that are currently visible on the page
			const visibleHeadings: IntersectionObserverEntry[] = [];
			Object.keys(headingElementsRef.current).forEach((key) => {
				const headingElement = headingElementsRef.current[key];
				if (headingElement.isIntersecting) visibleHeadings.push(headingElement);
			});

			const getIndexFromId = (id: any) =>
				headingElements.findIndex((heading) => heading.id === id);

			// If there is only one visible heading, this is our "active" heading
			if (visibleHeadings.length === 1) {
				setActiveId(visibleHeadings[0].target.id);
				// If there is more than one visible heading,
				// choose the one that is closest to the top of the page
			} else if (visibleHeadings.length > 1) {
				const sortedVisibleHeadings = visibleHeadings.sort(
					(a, b) => getIndexFromId(a.target.id) - getIndexFromId(b.target.id)
				);

				setActiveId(sortedVisibleHeadings[0].target.id);
			}
		};

		const observer = new IntersectionObserver(callback, {
			root: document.querySelector('iframe'),
			rootMargin: '500px',
		});

		const headingElements = Array.from(document.querySelectorAll('h2, h3'));

		headingElements.forEach((element) => observer.observe(element));

		return () => observer.disconnect();
	}, [setActiveId]);
};

/**
 * Renders the table of contents.
 */
const TableOfContents = () => {
	const [activeId, setActiveId] = useState<string>();
	const [headings, setHeadings] = useState<HTMLElement[]>([]);
	useEffect(() => {
		const headingElements: HTMLElement[] = Array.from(
			document.querySelectorAll('h2')
		);
		console.log(headingElements);
		setHeadings(headingElements);
	}, []);
	useIntersectionObserver(setActiveId);

	return (
		<aside className={styles.tocContainer}>
			<p>
				<strong>On This Page</strong>
			</p>
			<nav className={styles.toc} aria-label='Table of Contents'>
				<ul>
					{headings.map((heading: any) => (
						<li
							key={heading.id}
							className={heading.id === activeId ? `${styles.active}` : ''}
						>
							<a href={`#${heading.id}`}>{heading.innerText}</a>
						</li>
					))}
				</ul>
			</nav>
		</aside>
	);
};

export default TableOfContents;
