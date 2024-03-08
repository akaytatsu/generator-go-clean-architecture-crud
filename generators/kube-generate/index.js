import Generator from "yeoman-generator";

export default class extends Generator {

    initializing() {
        this.props = {};
    }

    async prompting() {
        this.answers = await this.prompt([{
            type: 'input',
            name: 'namespace',
            message: 'name of project like "vert-fileexplorer-prd", "vert-issue-prd", "vert-portal-hml": ',
        }]);
    }

    writing() {

        var namespace = this.answers.namespace;
        var basePath = './' + namespace + '/';

        var params = {
            namespace: namespace
        }

        this.fs.copyTpl(
            this.templatePath('namespace.yaml'),
            this.destinationPath(basePath + 'namespace.yaml'),
            params
        );

        this.fs.copyTpl(
            this.templatePath('configmap.yaml'),
            this.destinationPath(basePath + 'configmap.yaml'),
            params
        );

        this.fs.copyTpl(
            this.templatePath('configmap-front.yaml'),
            this.destinationPath(basePath + 'configmap-front.yaml'),
            params
        );

        this.fs.copyTpl(
            this.templatePath('frontend/frontend.yaml'),
            this.destinationPath(basePath + 'frontend/frontend.yaml'),
            params
        );

        this.fs.copyTpl(
            this.templatePath('frontend/new_ingress.yaml'),
            this.destinationPath(basePath + 'frontend/new_ingress.yaml'),
            params
        );

        this.fs.copyTpl(
            this.templatePath('frontend/service.yaml'),
            this.destinationPath(basePath + 'frontend/service.yaml'),
            params
        );

        this.fs.copyTpl(
            this.templatePath('backend/back-end.yaml'),
            this.destinationPath(basePath + 'backend/back-end.yaml'),
            params
        );

        this.fs.copyTpl(
            this.templatePath('backend/new_ingress.yaml'),
            this.destinationPath(basePath + 'backend/new_ingress.yaml'),
            params
        );

        this.fs.copyTpl(
            this.templatePath('backend/service.yaml'),
            this.destinationPath(basePath + 'backend/service.yaml'),
            params
        );
    }
}