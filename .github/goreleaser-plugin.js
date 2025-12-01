import { writeFile } from 'fs/promises';
import { execa } from 'execa';

/**
 * Custom semantic-release plugin to safely invoke goreleaser with release notes.
 * This avoids shell injection vulnerabilities by handling the release notes
 * directly in JavaScript rather than interpolating them into shell commands.
 */

export async function publish(pluginConfig, context) {
  const { nextRelease, logger } = context;
  const notesFile = '/tmp/release-notes.md';

  try {
    // Safely write release notes to file without shell interpolation
    logger.log('Writing release notes to %s', notesFile);
    await writeFile(notesFile, nextRelease.notes || '', 'utf-8');

    // Execute goreleaser with direct argument passing (no shell)
    logger.log('Running goreleaser release...');
    const result = await execa('goreleaser', [
      'release',
      '--release-notes',
      notesFile,
      '--clean'
    ], {
      stdio: 'inherit',
      preferLocal: true
    });

    logger.log('goreleaser completed successfully');
    return result;
  } catch (error) {
    logger.error('goreleaser failed: %s', error.message);
    throw error;
  }
}
